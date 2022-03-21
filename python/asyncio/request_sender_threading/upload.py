"""
upload.py

Given an NS1 request file (one NS1 request per line), uploads all objects to
the API at a given rate.
"""
import argparse
import logging
import queue
import os
import sys
import threading
from time import sleep
from timeit import default_timer
from typing import List, TextIO

import requests
from requests.packages.urllib3.exceptions import InsecureRequestWarning

import progress


RUNNING_IN_CONTAINER = ".dockerenv" in os.listdir("/")


def thread_boilerplate(targ, args):
    thread = threading.Thread(target=targ, args=args)
    thread.setDaemon(True)
    thread.start()
    logging.debug(f"Started helper thread {targ} with arguments {args}")
    return thread


class SharedCounter(object):
    """
    A counter singleton. Keeps tracks of values and is shared by all running
    threads.

    Tracks:
    * succesful requests in `self.success`
    * failed requests in `self.fail`
    * combined requests in `self.val`
    """

    _instance = None

    def __new__(cls, *args, **kwargs):
        if cls._instance is None:
            cls._instance = object.__new__(cls, *args, **kwargs)
        return cls._instance

    def __init__(self, value=0):
        self.fail = value
        self.success = value
        self.lock = threading.Lock()

    def incr_fail(self):
        """
        Increments the count of failed requests
        """
        with self.lock:
            self.fail += 1

    def incr_success(self):
        """
        Increments the count of succesful requests
        """
        with self.lock:
            self.success += 1

    @property
    def val(self):
        return self.fail + self.success


class Uploader(object):
    """
    The Uploader object - the worker that uploads to the API. Many of these may
    spawned for concurrent uploading.
    """

    def __init__(self, auth: str, verify: bool, base_url: str, objects: queue.Queue):
        self.sess = requests.Session()
        # self.sess.headers.update({"X-NSONE-Key": auth})
        self.sess.verify = verify
        self.base_url = base_url
        self.q = objects
        self.request_counter = SharedCounter()

    def run(self):
        while True:
            obj = self.q.get()

            if obj is None:
                break
            method, route, body = obj
            endpoint = f"{route}"

            try:
                max_tries = 10
                num_tries = 0

                while num_tries < max_tries:
                    response = self.sess.request(method.upper(), endpoint, data=body)
                    status = response.status_code
                    logging.debug(response.content)
                    if status == 429 or status > 499:
                        # Too Many Requests or higher than max (?)
                        logging.info(f"RETRYING {obj}: - {response.text.strip()}")
                        self.request_counter.incr_fail()
                        num_tries += 1
                        sleep(num_tries ** 2)
                    else:
                        break

                response.raise_for_status()
            except KeyboardInterrupt:
                raise
            except Exception as e:
                self.request_counter.incr_fail()
                logging.warning(
                    f"WARN - {obj} call could not be completed: {e} - {response.text.strip()}"
                )
            else:
                self.request_counter.incr_success()

            self.q.task_done()
        logging.debug("Worker thread complete")


def rate_limited_queue(ns1objs: List[str], q: queue.Queue, rate: int, fname: str) -> None:
    """
    Fills the queue object to be used by all workers. Tries to do so in a rate
    limited manner.
    """
    # FIXME rate limiting only vaguely works
    logging.debug(f"Starting to fill the queue, targeting {rate} rps")
    req_counter = SharedCounter()
    total = len(ns1objs)
    start_time = default_timer()
    t1 = default_timer()
    t3 = default_timer()
    prev_val = 0
    prev_suc = 0
    for i, ns1obj in enumerate(ns1objs):
        # a request line is structured:
        # method\x00endpoint\x00body
        # For example: 'PUT\x00v1/zones/frazao.ca\x00{"zone":"frazao.ca"}\n'
        # would be a PUT request for frazao.ca
        req = ns1obj.split("\x00")
        if len(req) != 3:
            logging.error(
                f"Unable to properly decode request on line {i+1}: {[char for char in ns1obj]}"
            )
            continue
        q.put(req)

        if i % rate == 0 and i > 0:
            while not q.empty():
                # If the queue is not empty we need to wait, otherwise we may feed
                # too much and lose control of the rate if the workers catch up
                logging.debug("sleeping, queue not empty")
                sleep(0.1)
            t2 = default_timer()
            sleep_time = 1 - (t2 - t1)
            t1 = default_timer()
            if sleep_time > 0:
                # If it took longer than one second for the worker to consume
                # the tasks then we dont need to sleep
                logging.debug(f"sleeping {sleep_time}, waiting for rate")
                sleep(sleep_time)

            if i % (rate * 10) == 0:
                elapsed_time = default_timer() - start_time
                cur_val = req_counter.val
                suc_val = req_counter.success
                cur_time = default_timer()
                cur_rate = (cur_val - prev_val) / (cur_time - t3)
                suc_rate = (suc_val - prev_suc) / (cur_time - t3)
                prev_val = cur_val
                prev_suc = suc_val
                t3 = cur_time
                logging.info(
                    "{"
                    + f'"rate":{cur_rate:.2f},"success_rate":{suc_rate:.2f},"total_requests":{cur_val},"success":{suc_val},"elapsed_time":{elapsed_time:.0f}'
                    + "}"
                )

            progress.progress(req_counter.success, req_counter.fail, total, fname)

    progress.progress(total, req_counter.fail, total, fname)
    if not RUNNING_IN_CONTAINER:
        print()  # Need to print a new line when were done printing the progress bar
    logging.info("Done filling queue...")


if __name__ == "__main__":
    help_texts = {
        "main": __doc__,
        "file": "Location of NS1 request file",
        "objtype": "Type of object to be ignest {zones, records, networks, subnets}",
        "apikey": "API key to make requests with",
        "rate": "Requests per second (default 5)",
        "api_host": "Base URL for API",
        "verify": "Turn off SSL verification",
        "debug": "Enable debug statements",
        "output_dir": "Location to store any output files",
    }

    parser = argparse.ArgumentParser(
        description=help_texts["main"], formatter_class=argparse.RawTextHelpFormatter
    )

    parser.add_argument("file", type=str, help=help_texts["file"])
    parser.add_argument("apikey", type=str, help=help_texts["apikey"])

    parser.add_argument(
        "-a",
        "--api_host",
        type=str,
        default="localhost",
        help=help_texts["api_host"],
    )

    parser.add_argument(
        "-r",
        "--rate",
        default=5,
        type=int,
        help=help_texts["rate"],
    )

    parser.add_argument(
        "-v",
        "--verify",
        action="store_false",
        help=help_texts["verify"],
    )

    parser.add_argument(
        "-d",
        "--debug",
        action="store_true",
        help=help_texts["debug"],
    )

    parser.add_argument(
        "-o",
        "--output_dir",
        type=str,
        default="ns1",
        help=help_texts["output_dir"],
    )


    args = parser.parse_args()
    logging.basicConfig(
        level=logging.DEBUG,
        format="%(asctime)s - %(thread)d - %(levelname)s - %(filename)s: %(message)s",
        handlers=[
            logging.StreamHandler()
            if RUNNING_IN_CONTAINER
            else logging.FileHandler(os.path.join(args.output_dir, "alchemist.log"))
        ],
    )

    rate = args.rate
    if not args.verify:
        requests.packages.urllib3.disable_warnings(InsecureRequestWarning)

    if not args.debug:
        logging.disable(logging.DEBUG)

    logging.info(f"Start of Program - using file {args.file}")
    req_counter = SharedCounter()

    work_q: queue.Queue = queue.Queue()

    num_workers = 100
    threads = []
    logging.info(f"Using {num_workers} threads")

    objects = [line.strip() for line in open(args.file).readlines() if line.strip()]
    q_fill_thread = thread_boilerplate(rate_limited_queue, (objects, work_q, rate, args.file))

    for i in range(num_workers):
        thread = threading.Thread(
            target=Uploader(
                args.apikey, args.verify, f"https://{args.api_host}", work_q
            ).run,
            args=(),
        )
        threads.append(thread)
        thread.setDaemon(True)
        thread.start()

        logging.debug(f"Starting worker thread: {i}")

    work_q.join()

    q_fill_thread.join()
    for i in range(num_workers):
        work_q.put(None)

    for thread in threads:
        thread.join()

    logging.info(
        f"End of Program - completed {req_counter.val} requests - {req_counter.success} succesful"
    )

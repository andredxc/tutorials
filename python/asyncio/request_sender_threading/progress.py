import os
import sys
import shutil
import math
import logging
from timeit import default_timer


RUNNING_IN_CONTAINER = ".dockerenv" in os.listdir("/")


start = default_timer()
cur_time = start + 1
prev_done = 0
last_log = 0
# Interval to log progress if running in container
INTERVAL = 60


def progress(done, errors, total, fname):
    global cur_time, prev_done, last_log

    bar_size = 50
    pct_done = done / total if total > 0 else 1
    done_bars = math.ceil(pct_done * bar_size)

    cur_rate = (done - prev_done) / (default_timer() - cur_time)
    prev_done = done
    cur_time = default_timer()
    avg_rate = done / (default_timer() - start)
    if done == total:
        cur_rate = avg_rate
    eta = int((total - done) / avg_rate) if done else 0
    elapsed = int(cur_time - start)

    if RUNNING_IN_CONTAINER:
        if (cur_time - last_log) > INTERVAL:
            last_log = default_timer()
            logging.info(
                f"{fname} {done:,d} / {total:,d} ({pct_done*100:0.0f}%) - Avg Rate {avg_rate:,.0f} - Errors: {errors:,d} - ELAPSED: {elapsed:,d}s - ETA: {eta:,d}s"
            )
    else:
        print_str = "".join(
            [
                f"\r{fname[:30]} [",
                "#" * done_bars,
                " " * (bar_size - done_bars),
                f"] {done:,d} / {total:,d} ({pct_done*100:0.0f}%) - "
                f"Avg Rate {avg_rate:,.0f} - Errors {errors:,d} - "
                f"ELAPSED: {elapsed:,d}s - ETA: {eta:,d}s",
            ]
        )
        print_str = print_str.ljust(shutil.get_terminal_size()[0])
        print(print_str, file=sys.stderr, end="")
        sys.stderr.flush()

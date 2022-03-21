#!/usr/bin/python3.8
import asyncio

async def say_after(seconds, text):
	await asyncio.sleep(seconds)
	print(text)

async def wait_1s():
	await asyncio.sleep(1)

async def main():

	task1 = asyncio.create_task(say_after(1, "One"))
	task2 = asyncio.create_task(say_after(2, "Two"))
	task3 = asyncio.create_task(say_after(3, "Three"))
	task4 = asyncio.create_task(say_after(4, "Go!"))

	print('Waiting')
	await task1
	await task2
	await task3
	await task4
	
	task5 = asyncio.create_task(wait_1s())
	await task5

	print('Done')

asyncio.run(main())
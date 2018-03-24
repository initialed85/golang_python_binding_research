import time
from multiprocessing import Process

from src.attempt5 import DoSomething


def handle_do_something():
    for i in range(0, 8):
        DoSomething()


# won't work- can't call the module from the parent, seems to break everything
# print 'out of process'
# handle_do_something()

print 'in process'
processes = [Process(target=handle_do_something) for _ in range(0, 8)]

for p in processes:
    p.start()

time.sleep(2)

for p in processes:
    p.join()

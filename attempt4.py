import datetime
import time
from multiprocessing import Process

from src.attempt4 import DoSomething

print 'Doing something in a process'
p = Process(target=DoSomething)
p.start()
time.sleep(2)
p.join()

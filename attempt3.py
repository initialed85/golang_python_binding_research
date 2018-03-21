import datetime

from threading import Thread

from src.attempt3 import InefficientWork, EfficientWork

threads = []
for i in range(0, 10):
    t = Thread(
        target=InefficientWork
    )
    threads += [t]

before = datetime.datetime.now()
for t in threads:
    t.start()

for t in threads:
    t.join()
after = datetime.datetime.now()

print 'InefficientWork took', after - before

del threads

threads = []
for i in range(0, 10):
    t = Thread(
        target=EfficientWork
    )
    threads += [t]

before = datetime.datetime.now()
for t in threads:
    t.start()

for t in threads:
    t.join()
after = datetime.datetime.now()

print 'EfficientWork took', after - before

from sys import version as python_version

if 'pypy' not in python_version.strip().lower():
    from py2.attempt2 import *
else:
    from cffi.attempt2 import *

    SetPyPy()

    print 'WARNING: PyPy detected- be prepared for very odd behaviour'

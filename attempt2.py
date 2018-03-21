from src.attempt2 import GetSomeStruct


for i in range(0, 1000000):
    some_struct = GetSomeStruct()
    print i, len(some_struct.SomeString)
    del (some_struct)

This purpose of this attempt is to test GIL behaviour.

You'll notice that with Python, InefficientWork cannot complete concurrently (because of the GIL on the Python side).

EfficientWork releases the GIL so it's all good.

For PyPy (using CFFI) there doesn't seem to be the same problem; presumably CFFI deals with the GIL in it's own way.
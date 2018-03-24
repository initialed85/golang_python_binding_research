from importlib import import_module
from sys import version as python_version


class __GoImporter(object):
    def __init__(self, module_name, obj_names, python_module_path=None, pypy_module_path=None):
        self.pypy = 'pypy' in python_version.strip().lower()
        self.python_module_path = python_module_path if python_module_path is not None else 'py2'
        self.pypy_module_path = pypy_module_path if pypy_module_path is not None else 'cffi'

        self.module_path = '{0}.{1}.{2}'.format(
            __name__,
            self.pypy_module_path if self.pypy else self.python_module_path,
            module_name,
        )

        for obj_name in obj_names:
            globals()[obj_name] = self.__build_wrapper(obj_name)

    def __build_wrapper(self, obj_name):
        def wrapper(*args, **kwargs):
            module = import_module(self.module_path)

            if self.pypy:
                print 'WARNING: PyPy detected- be prepared for very odd behaviour'
                try:
                    module.SetPyPy()
                except BaseException:
                    pass

            obj = getattr(
                module,
                obj_name
            )

            return obj(*args, **kwargs)

        return wrapper


__go_importer = __GoImporter(
    module_name='attempt5',
    obj_names=['SetPyPy', 'DoSomething']
)

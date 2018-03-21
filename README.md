## golang_python_binding_research

This repo contains some of my research and testing for accessing Go libraries from Python

#### Prerequisites

* Go1.9 or Go1.10
* Python2.7 and PyPy5.10
* virtualenvwrapper
* pkgconfig/pkg-config

#### Setup

* ```mkvirtualenv -p `which python2.7` golang_python_binding_research_py2```
* ```mkvirtualenv -p `which pypy` golang_python_binding_research_pypy```

#### Build
* ```./build.sh```

#### Test
* Terminal window 1
    * ```./build.sh```
* Terminal window 2
    * ```workon golang_python_binding_research_py2```
    * ```python attempt(n).py```
* Terminal window 3
    * ```workon golang_python_binding_research_pypy```
    * ```python attempt(n).py```

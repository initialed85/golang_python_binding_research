The purpose of this attempt is to see how things deal with multiprocessing (fork() I guess?)

On MacOS at least, it doesn't work- you end up with a big old stack trace. Interestingly, it seems to run some of the method on the Go side.

Stack traces below:

**** Python

(golang_python_binding_research_py2) theophilus:golang_python_binding_research edwardbeech$ clear; GODEBUG=cgocheck=0 python attempt4.py
Doing something in a process
before doing something
mach error semaphore_signal: 15
fatal error: mach error

runtime stack:
runtime.throw(0x105d7c921, 0xa)
	/usr/local/Cellar/go/1.10/libexec/src/runtime/panic.go:619 +0x81
runtime.macherror(0x10000000f, 0x105d7d61e, 0x10)
	/usr/local/Cellar/go/1.10/libexec/src/runtime/os_darwin.go:224 +0xa3
runtime.mach_semrelease.func1()
	/usr/local/Cellar/go/1.10/libexec/src/runtime/os_darwin.go:477 +0x3d
runtime.mach_semrelease(0x7ffe00001603)
	/usr/local/Cellar/go/1.10/libexec/src/runtime/os_darwin.go:477 +0x45
runtime.semawakeup(0xc420046800)
	/usr/local/Cellar/go/1.10/libexec/src/runtime/os_darwin.go:36 +0x21
runtime.notewakeup(0xc420046948)
	/usr/local/Cellar/go/1.10/libexec/src/runtime/lock_sema.go:147 +0x59
runtime.startm(0x0, 0xc420000101)
	/usr/local/Cellar/go/1.10/libexec/src/runtime/proc.go:2012 +0xda
runtime.wakep()
	/usr/local/Cellar/go/1.10/libexec/src/runtime/proc.go:2078 +0x48
runtime.newproc1(0x105d83090, 0xc420054d58, 0x100000008, 0x105cf57a7)
	/usr/local/Cellar/go/1.10/libexec/src/runtime/proc.go:3343 +0x314
runtime.newproc.func1()
	/usr/local/Cellar/go/1.10/libexec/src/runtime/proc.go:3239 +0x46
runtime.systemstack(0x7ffeea39a6c0)
	/usr/local/Cellar/go/1.10/libexec/src/runtime/asm_amd64.s:409 +0x79
runtime.mstart()
	/usr/local/Cellar/go/1.10/libexec/src/runtime/proc.go:1170

goroutine 17 [running, locked to thread]:
runtime.systemstack_switch()
	/usr/local/Cellar/go/1.10/libexec/src/runtime/asm_amd64.s:363 fp=0xc420054d08 sp=0xc420054d00 pc=0x105d00f90
runtime.newproc(0xc400000008, 0x105d83090)
	/usr/local/Cellar/go/1.10/libexec/src/runtime/proc.go:3238 +0x64 fp=0xc420054d48 sp=0xc420054d08 pc=0x105ce3f74
runtime.(*timersBucket).addtimerLocked(0x105dff460, 0xc420020100)
	/usr/local/Cellar/go/1.10/libexec/src/runtime/time.go:160 +0x107 fp=0xc420054da0 sp=0xc420054d48 pc=0x105cf57a7
time.Sleep(0x3b9aca00)
	/usr/local/Cellar/go/1.10/libexec/src/runtime/time.go:101 +0x135 fp=0xc420054e00 sp=0xc420054da0 pc=0x105cf5535
attempt4.DoSomething(0x0, 0xc420054e88)
	/Users/edwardbeech/Projects/Home/golang_python_binding_research/src/attempt4/attempt4.go:17 +0x6e fp=0xc420054e60 sp=0xc420054e00 pc=0x105d46c2e
main.cgo_func_attempt4_DoSomething(0xc400000008, 0x105d82c80)
	/private/var/folders/j1/b2w166dj3hx8fn68q1ys0r740000gn/T/gopy-209737331/attempt4.go:167 +0x22 fp=0xc420054e80 sp=0xc420054e60 pc=0x105d48e12
main._cgoexpwrap_0d8f056a40e9_cgo_func_attempt4_DoSomething(0x0, 0x0)
	_cgo_gotypes.go:219 +0x50 fp=0xc420054ea8 sp=0xc420054e80 pc=0x105d48460
runtime.call32(0x0, 0x7ffeea39a710, 0x7ffeea39a7a8, 0x10)
	/usr/local/Cellar/go/1.10/libexec/src/runtime/asm_amd64.s:573 +0x3b fp=0xc420054ed8 sp=0xc420054ea8 pc=0x105d0135b
runtime.cgocallbackg1(0x0)
	/usr/local/Cellar/go/1.10/libexec/src/runtime/cgocall.go:316 +0x19c fp=0xc420054f58 sp=0xc420054ed8 pc=0x105cb57bc
runtime.cgocallbackg(0x0)
	/usr/local/Cellar/go/1.10/libexec/src/runtime/cgocall.go:194 +0xda fp=0xc420054fc0 sp=0xc420054f58 pc=0x105cb558a
runtime.cgocallback_gofunc(0x0, 0x0, 0x0, 0x0)
	/usr/local/Cellar/go/1.10/libexec/src/runtime/asm_amd64.s:826 +0x9b fp=0xc420054fe0 sp=0xc420054fc0 pc=0x105d0292b
runtime.goexit()
	/usr/local/Cellar/go/1.10/libexec/src/runtime/asm_amd64.s:2361 +0x1 fp=0xc420054fe8 sp=0xc420054fe0 pc=0x105d03a61
(golang_python_binding_research_py2) theophilus:golang_python_binding_research edwardbeech$

**** PyPy

(golang_python_binding_research_pypy) theophilus:golang_python_binding_research edwardbeech$ clear; GODEBUG=cgocheck=0 python attempt4.py
WARNING: PyPy detected- be prepared for very odd behaviour
Doing something in a process
before doing something
mach error semaphore_signal: 15
fatal error: mach error

runtime stack:
runtime.throw(0x110c2f901, 0xa)
	/usr/local/Cellar/go/1.10/libexec/src/runtime/panic.go:619 +0x81
runtime.macherror(0x10000000f, 0x110c305fe, 0x10)
	/usr/local/Cellar/go/1.10/libexec/src/runtime/os_darwin.go:224 +0xa3
runtime.mach_semrelease.func1()
	/usr/local/Cellar/go/1.10/libexec/src/runtime/os_darwin.go:477 +0x3d
runtime.mach_semrelease(0x7ffe00001603)
	/usr/local/Cellar/go/1.10/libexec/src/runtime/os_darwin.go:477 +0x45
runtime.semawakeup(0xc420046800)
	/usr/local/Cellar/go/1.10/libexec/src/runtime/os_darwin.go:36 +0x21
runtime.notewakeup(0xc420046948)
	/usr/local/Cellar/go/1.10/libexec/src/runtime/lock_sema.go:147 +0x59
runtime.startm(0x0, 0xc420000101)
	/usr/local/Cellar/go/1.10/libexec/src/runtime/proc.go:2012 +0xda
runtime.wakep()
	/usr/local/Cellar/go/1.10/libexec/src/runtime/proc.go:2078 +0x48
runtime.newproc1(0x110c36070, 0xc420054d58, 0x100000008, 0x110ba88f7)
	/usr/local/Cellar/go/1.10/libexec/src/runtime/proc.go:3343 +0x314
runtime.newproc.func1()
	/usr/local/Cellar/go/1.10/libexec/src/runtime/proc.go:3239 +0x46
runtime.systemstack(0x7ffee269fd90)
	/usr/local/Cellar/go/1.10/libexec/src/runtime/asm_amd64.s:409 +0x79
runtime.mstart()
	/usr/local/Cellar/go/1.10/libexec/src/runtime/proc.go:1170

goroutine 17 [running, locked to thread]:
runtime.systemstack_switch()
	/usr/local/Cellar/go/1.10/libexec/src/runtime/asm_amd64.s:363 fp=0xc420054d08 sp=0xc420054d00 pc=0x110bb40e0
runtime.newproc(0xc400000008, 0x110c36070)
	/usr/local/Cellar/go/1.10/libexec/src/runtime/proc.go:3238 +0x64 fp=0xc420054d48 sp=0xc420054d08 pc=0x110b970c4
runtime.(*timersBucket).addtimerLocked(0x110cb23e0, 0xc420020100)
	/usr/local/Cellar/go/1.10/libexec/src/runtime/time.go:160 +0x107 fp=0xc420054da0 sp=0xc420054d48 pc=0x110ba88f7
time.Sleep(0x3b9aca00)
	/usr/local/Cellar/go/1.10/libexec/src/runtime/time.go:101 +0x135 fp=0xc420054e00 sp=0xc420054da0 pc=0x110ba8685
attempt4.DoSomething(0x0, 0xc420054e88)
	/Users/edwardbeech/Projects/Home/golang_python_binding_research/src/attempt4/attempt4.go:17 +0x6e fp=0xc420054e60 sp=0xc420054e00 pc=0x110bf9d7e
main.cgo_func_attempt4_DoSomething(0xc400000008, 0x110c35c60)
	/private/var/folders/j1/b2w166dj3hx8fn68q1ys0r740000gn/T/gopy-299733581/attempt4.go:167 +0x22 fp=0xc420054e80 sp=0xc420054e60 pc=0x110bfbf62
main._cgoexpwrap_3134bae50931_cgo_func_attempt4_DoSomething(0x0, 0x0)
	_cgo_gotypes.go:219 +0x50 fp=0xc420054ea8 sp=0xc420054e80 pc=0x110bfb5b0
runtime.call32(0x0, 0x7ffee269fde0, 0x7ffee269fe78, 0x10)
	/usr/local/Cellar/go/1.10/libexec/src/runtime/asm_amd64.s:573 +0x3b fp=0xc420054ed8 sp=0xc420054ea8 pc=0x110bb44ab
runtime.cgocallbackg1(0x0)
	/usr/local/Cellar/go/1.10/libexec/src/runtime/cgocall.go:316 +0x19c fp=0xc420054f58 sp=0xc420054ed8 pc=0x110b6890c
runtime.cgocallbackg(0x0)
	/usr/local/Cellar/go/1.10/libexec/src/runtime/cgocall.go:194 +0xda fp=0xc420054fc0 sp=0xc420054f58 pc=0x110b686da
runtime.cgocallback_gofunc(0x0, 0x0, 0x0, 0x0)
	/usr/local/Cellar/go/1.10/libexec/src/runtime/asm_amd64.s:826 +0x9b fp=0xc420054fe0 sp=0xc420054fc0 pc=0x110bb5a7b
runtime.goexit()
	/usr/local/Cellar/go/1.10/libexec/src/runtime/asm_amd64.s:2361 +0x1 fp=0xc420054fe8 sp=0xc420054fe0 pc=0x110bb6bb1
(golang_python_binding_research_pypy) theophilus:golang_python_binding_research edwardbeech$
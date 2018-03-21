package common

// #cgo pkg-config: python2
// #include <Python.h>
import "C"

var runningPyPy = false

func releaseGIL() *C.PyThreadState {
	if runningPyPy {
		return nil
	}

	var tState *C.PyThreadState

	tState = C.PyEval_SaveThread()

	return tState
}

func reacquireGIL(tState *C.PyThreadState) {
	if tState == nil {
		return
	}

	C.PyEval_RestoreThread(tState)
}

// SetPyPy is used by the Python side to declare whether or not we're running under PyPy (not easily discovered on the Go side)
func SetPyPy() {
	runningPyPy = true
}

// GetPyPy returns true if we're running under PyPy
func GetPyPy() bool {
	val := runningPyPy

	return val
}

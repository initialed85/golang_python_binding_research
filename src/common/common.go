package common

// #cgo pkg-config: python2
// #include <Python.h>
import "C"

var runningPyPy = false

// SetPyPy is used by the Python side to declare whether or not we're running under PyPy
func SetPyPy() {
	runningPyPy = true // no mutex around this for speed
}

// ReleaseGIL releases the Python GIL
func ReleaseGIL() *C.PyThreadState {
	if runningPyPy {
		return nil
	}

	var tState *C.PyThreadState

	tState = C.PyEval_SaveThread()

	return tState
}

// ReacquireGIL reacquires a previously released Python GIL
func ReacquireGIL(tState *C.PyThreadState) {
	if tState == nil {
		return
	}

	C.PyEval_RestoreThread(tState)
}

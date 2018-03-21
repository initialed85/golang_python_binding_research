package attempt3

import (
	"common"
	"time"
)

// SetPyPy is used by the Python side
func SetPyPy() {
	common.SetPyPy()
}

// InefficientWork sleeps for a second without touching the GIL
func InefficientWork() {
	time.Sleep(time.Second)
}

// EfficientWork sleeps for a second with the GIL released
func EfficientWork() {
	threadState := common.ReleaseGIL()
	time.Sleep(time.Second)
	common.ReacquireGIL(threadState)
}

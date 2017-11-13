package attempt_6

// #cgo pkg-config: python
// #include <Python.h>
import "C" // the lack of a space between this and the above is CRITICAL

import (
	"fmt"
	"time"
)

func BackgroundWork() {
	tState := C.PyEval_SaveThread()

	for i := 0; i < 10; i++ {
		fmt.Println("go: work")
		time.Sleep(time.Second * 1)
	}

	C.PyEval_RestoreThread(tState)
}

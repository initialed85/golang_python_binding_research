package attempt4

import (
	"common"
	"fmt"
	"time"
)

// SetPyPy is used by the Python side
func SetPyPy() {
	common.SetPyPy()
}

// DoSomething does something
func DoSomething() string {
	fmt.Println("Before doing something")
	time.Sleep(time.Second)
	fmt.Println("After doing something")

	return "Okay, I did it"
}

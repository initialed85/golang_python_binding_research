package attempt2

import (
	"bytes"
	"common"
)

type SomeStruct struct {
	SomeString string
}

// SetPyPy is used by the Python side
func SetPyPy() {
	common.SetPyPy()
}

// GetSomeStruct returns a struct with a big string in it
func GetSomeStruct() SomeStruct {
	var buf bytes.Buffer

	for i := 0; i < 10000000; i++ {
		buf.WriteString("a")
	}

	return SomeStruct{
		SomeString: buf.String(),
	}
}

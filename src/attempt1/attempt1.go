package attempt1

import "common"

type OneWithTheLot struct {
	SomeString        string
	SomeInt           int64
	SomeFloat         float64
	SomeBool          bool
	SomeListOfStrings []string
	SomeListOfInts    []int64
	SomeListOfFloats  []float64
	SomeListOfBools   []bool
}

// SetPyPy is used by the Python side
func SetPyPy() {
	common.SetPyPy()
}

// GetOneWithTheLot returns a struct with a public properties of different types
func GetOneWithTheLot() OneWithTheLot {
	return OneWithTheLot{
		"some string",
		1337,
		1337.1337,
		true,
		[]string{"some", "list", "of", "strings"},
		[]int64{6, 2, 9, 1},
		[]float64{6.6, 2.2, 9.9, 1.1},
		[]bool{true, false, true, false},
	}
}

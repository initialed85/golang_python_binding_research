package sum

import "errors"

type SomeObject struct {
	Name string
}

func NewSomeObject(name string) SomeObject {
	someObject := SomeObject{
		Name: name,
	}

	return someObject
}

func Sum(a, b int) (float64, error) {
	if a == b {
		return 0.0, errors.New("oh bugger")
	}

	return float64(a) + float64(b), nil
}


package sum

import "C"

//export Sum
func Sum(a, b int) int {
	return a + b
}

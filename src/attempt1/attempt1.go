package attempt1

import "fmt"

// SayHelloTo returns a greeting string for the given name
func SayHelloTo(name string) string {
	return fmt.Sprintf("Hello, %s", name)
}

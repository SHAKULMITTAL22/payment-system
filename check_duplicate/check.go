// main.go
package main

import (
	"fmt"
)

// ExportedFunction is the function that will be exported.
func InternalFunction() {
	result := internalFunction("Hello from the internal function!")
	fmt.Println(result)
}

// internalFunction is an unexported (internal) helper function.
func internalFunction(message string) string {
	return fmt.Sprintf("Internal Function Received: %s", message)
}

func main() {
	// Call the exported function.
	InternalFunction()
}

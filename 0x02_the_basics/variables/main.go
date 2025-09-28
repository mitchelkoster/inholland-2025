/*
Learn about variables, data types and the walrus operator.
*/
package main

import "fmt"

var globalMessage string = "Global message: initial"

func main() {
	var name string  // Declared, not initalized
	name = "Mitchel" // Initialized

	var age int = 36 // Declared and initialized

	// Declared, initialized, but not used!
	// var money float32 = 0.0

	// Printing multple values
	fmt.Println("Hello,", name, age)

	// Formatted prints look good!
	fmt.Printf("Hello, %s (%T) %d (%T)\n", name, name, age, age)

	var a, b, c int = 1, 2, 3
	fmt.Println(a, b, c)

	// The := syntax is shorthand for declaring and
	// initializing a variable
	message := "Golang is kinda cool!"
	fmt.Println(message)

	// Const, cannot be  overwritten (high precision calculation)
	const VERSION = "v1.0.0"
	// VERSION = "v1.0.0" // Errors
	fmt.Println("Program version: ", VERSION)

	// Overwrites globalmessage
	// globalMessage = "Gobal message: updated"

	// Creates a new global variable, but only in the
	// scope of this function
	var globalMessage string = "Global mesage: scope"
	fmt.Println(globalMessage)

	test()
}

func test() {
	fmt.Println(globalMessage)
}

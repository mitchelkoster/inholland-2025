/*
Learn about variables, data types and the walrus operator.
*/
package main

import "fmt"

func main() {
	var name string  // Declared, not initalized
	name = "Mitchel" // Initialized

	var age int = 36 // Declared and initialized

	// Declared, initialized, but not used!
	// var money float32 = 0.0

	// Printing multple values
	fmt.Println("Hello,", name, age)

	// Formatted prints look good!
	fmt.Printf("Hello, %s (age %d)!\n", name, age)

	var a, b, c int = 1, 2, 3
	fmt.Println(a, b, c)

	// The := syntax is shorthand for declaring and
	// initializing a
	message := "Golang is kinda cool!"
	fmt.Println(message)
}

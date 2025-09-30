package main

import (
	"fmt"
	"runtime"
	"time"
)

// basic function
func sumNumber2(a int, b int) int {
	return a + b
}

func main() {
	isTrue := true
	isFalse := false

	// Defer to block end
	defer fmt.Println("Run at the end!")

	// True even if the first case matches
	if isTrue || isFalse == true {
		fmt.Println("I am true, even if the first statement matches")
	}

	// If statements are simple
	if isTrue && isFalse {
		fmt.Println("Both true")
	} else if isTrue && !isFalse {
		fmt.Println("First true, second false!")
	} else {
		fmt.Println("Other")
	}

	// Swicthes execute in order and are not limited to integers
	// https://github.com/mitchelkoster/vacay-gaurdian/blob/1.0.0/utils/storage.go
	fmt.Print("The system is running....")
	switch os := runtime.GOOS; os {
	case "windows":
		fmt.Println("macOS")
	case "linux":
		fmt.Println("Linux")
	case "darwin":
		fmt.Println("Mac")
	default:
		fmt.Printf("%s.\n", os)
	}

	// You can do calculations in cases
	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}

	// Switches don't need conditions and can act
	// like long if-else chains
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}

	// Switches can have multiple conditions
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend")
	default:
		fmt.Println("It's a weekday")
	}

	// WHILE
	i := 0
	for i < 5 {
		fmt.Print(i)
		i++
	}
	fmt.Println()

	// ENDLESS
	// for {
	// 	fmt.Println("Hello world!")
	// }

	// FOR loop
	for i := 5; i < 10; i++ {
		fmt.Print(i)
	}
	fmt.Println()

	// Ranges (basic)
	for i := range 5 {
		fmt.Print(i)
	}
	fmt.Println()

	// Ranges apply to built-in types (foreach)
	nums := []int{2, 3, 4}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum)

	// Defers are stacked function calls (event loop-like)
	// It's essentially a "finally" exception block
	// Can help prevent deadlocks
	fmt.Println("counting")
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}
	fmt.Println("done")

	var sumNumber1 = func(a int, b int) int {
		return a + b
	}
	fmt.Println("Basic function: ", sumNumber1(1, 3))
	fmt.Println("Basic function: ", sumNumber2(2, 7))

	// fmt.Println("Basic function: ", sumNumber3(3, 10)) // not defined

	// "go run ./main" will only compile main -> fail
	// "go run ." will compile the "package", but no "go init ." -> fail
	// "go run ./main ./functions"
	fmt.Println("Basic function: ", SumNumber3(3, 10))

	// Original search
	name, age := FindUser("Mitchel")
	fmt.Printf("Name: %s, age: %d\n", name, age)

	// Named return search
	name, age = FindUser2("Mitchel")
	fmt.Printf("Name: %s, age: %d\n", name, age)

	// Variadic function
	FindUser3("Mitchel", "Elizabeth", "Esmée")

	// Anonymous functions
	func() {
		fmt.Println("Anonymous function call")
	}()
}

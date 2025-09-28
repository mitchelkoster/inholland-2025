/*
Learn arrays and basic array operations
*/

package main

import "fmt"

func main() {
	// Declaring a default array
	var arrayA [5]int

	fmt.Println("array A:", arrayA)

	// Append values to an array
	var arrayB [5]int
	arrayB[0] = 1
	arrayB[1] = 2
	arrayB[2] = 3
	arrayB[3] = 4
	arrayB[4] = 5

	fmt.Println("array B:", arrayB)

	// Declare and initialize with the walrus operator
	arrayC := [4]int{1, 9, 8, 8}

	fmt.Println("array C:", arrayC)

	// The compiler can count the elements for you
	arrayD := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println("array D:", arrayD)

	// But won´t always work when overwriting exiting variables
	// arrayC = [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	// 2+D arrays are also possible of course
	playerPosition := [5][2]int{
		{0, 3},
		{2, 9},
		{1, 0},
		{8, 2},
	}

	// Other array operations are normal
	fmt.Println("\nPlayer positions:", playerPosition)
	fmt.Printf(
		"Player 1 position X: %d, Y: %d\n",
		playerPosition[0][0], playerPosition[0][1])
	fmt.Printf(
		"Player 2 position X: %d, Y: %d\n",
		playerPosition[1][0],
		playerPosition[1][1])
}

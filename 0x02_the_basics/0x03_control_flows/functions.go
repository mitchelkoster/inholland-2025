package main

import (
	"fmt"
	"strconv"
)

// lower case = local scope
func sumNumber3(a int, b int) int {
	return a + b
}

// Funcition overriding does not exist
// func sumNumber3(a float, b float) float {
// 	return a + b
// }

// Uppercase is exported
func SumNumber3(a int, b int) int {
	return a + b
}

// We only have to declare a type once if it matches
func SumNumber4(a, b, c int) int {
	return a + b + c
}

// We can also have more than one return value
func FindUser(username string) (string, int) {
	users := [3][2]string{
		{"Mitchel", "36"},
		{"Elizabeth", "20 ;-)"},
		{"Esmée", "4"},
	}

	for _, val := range users {
		name := val[0]
		age := val[1]

		// We should really handle errors
		// "Elizabeth" 0
		i, _ := strconv.Atoi(age)

		if name == username {
			return name, i
		}
	}

	return "", 0
}

// Named return values (also declare at function call)
func FindUser2(username string) (name string, age int) {
	users := [3][2]string{
		{"Mitchel", "36"},
		{"Elizabeth", "20 ;-)"},
		{"Esmée", "4"},
	}

	for _, val := range users {
		name = val[0]
		tmpAge := val[1]

		// We should really handle errors
		// "Elizabeth" 0
		age, _ = strconv.Atoi(tmpAge)

		if name == username {
			return
		}
	}

	return
}

// Variadic function
// func FindUser3(username1 string, username2 string) (string, int) {
// func FindUser3(usernames [10]string) (string, int) {
func FindUser3(usernames ...string) {
	for key, val := range usernames {
		fmt.Println(key, "=>", val)
	}
}

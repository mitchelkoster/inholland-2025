/*
Show example of file handling.

https://github.com/mitchelkoster/vacay-gaurdian/blob/main/utils/storage.go
*/

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// READING FILES
	fh, err := os.Open("files/access.log")

	// See "go doc os.Open"
	if err != nil {
		fmt.Println("Error:", err)
		return // Early return
	}
	defer fh.Close() // Memory leak wihtout defer

	// Read file content
	scanner := bufio.NewScanner(fh)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	// Look for errors while reading
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	// Create a new file
	fh2, err := os.Create("files/test.txt")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer fh2.Close()

	// WRITING TO FILES
	_, err = fh2.WriteString("Hello world!" + "\n")
	if err != nil {
		fmt.Println("Error writing line:", err)
		return
	}

	// APPEND TO A FILE (more controll)
	fh3, err := os.OpenFile("files/test.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer fh.Close() // altijd sluiten

	_, err = fh3.WriteString("Goodbye world!" + "\n")
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}
}

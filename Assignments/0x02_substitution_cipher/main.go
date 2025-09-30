/*
This an excersize to master functions, and basic control flows
*/
package main

import (
	"fmt"
	"learning-go/cipher/custom_crypto"
)

const secretString = "Welcome to the InHolland 2025 Golang course! I hope you learn something new today!"

func main() {
	fmt.Println("Chracter Substitution Cipher v1.0")

	// Wrypbfr gb gur IaHbyynaq 2025 Gbynae pbhtzr! I ublr cbh yrnta zbfrguvae arx gbqnc!
	encoded := custom_crypto.EncodeCipher(secretString)
	decoded := custom_crypto.DecodeCipher(encoded)

	fmt.Println()
	fmt.Println("Encoded :", encoded)
	fmt.Println("Decoded :", decoded)
}

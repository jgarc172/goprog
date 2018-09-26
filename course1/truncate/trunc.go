package main

import (
	"fmt"
)

func main() {
	var anInt int
	var aFloat float32
	fmt.Print("Enter a decimal number: ")
	fmt.Scan(&aFloat)
	anInt = int(aFloat)
	fmt.Printf("Your number truncated as an integer: %v\n", anInt)
}

package main

import (
	"fmt"
)

func main() {
	var userNumber float32

	fmt.Println("Enter a floating point number")
	_, err := fmt.Scan(&userNumber)

	if err != nil {
		fmt.Println("Error occured", err)

		// prevent Scan from passing the string to stdout if wrong type
		var restString string
		fmt.Scanln(&restString)

	} else {
		fmt.Printf("%v\n", int32(userNumber))
	}
}

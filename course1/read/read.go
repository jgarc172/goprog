// Package read defines a struct type name
package main

import (
	"fmt"
)

func main() {
	n := name{"Joe", "Doe"}
	fmt.Println(n.fname, n.lname)
}

type name struct {
	fname string
	lname string
}

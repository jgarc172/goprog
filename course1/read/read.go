// Package main reads names from a file, adds them to a list and prints the list
package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fileName := readFileName()
	file, fErr := os.Open(fileName)
	if fErr != nil {
		fmt.Println("failed to open file", fileName)
		os.Exit(1)
	}
	defer file.Close()

	names := make([]name, 0, 3)
	var n name
	var err error
	for err != io.EOF {
		n, err = readName(file)
		if err == nil {
			names = append(names, n)
		}
	}
	printNames(names)
}

type name struct {
	fname string
	lname string
}

func readFileName() (fileName string) {
	fmt.Print("File name: ")
	fmt.Scanln(&fileName)
	return
}
func readName(r io.Reader) (name, error) {
	var n name
	var f, l string
	_, err := fmt.Fscanln(r, &f, &l)
	if err != nil {
		return n, io.EOF
	}
	n = name{f, l}
	return n, nil
}
func printNames(names []name) {
	for _, n := range names {
		fmt.Println(n.fname, n.lname)
	}
}

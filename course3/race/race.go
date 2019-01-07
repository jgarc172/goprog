package main

import "fmt"

var x = 0

func main() {
	go increment()
	go increment()
	fmt.Println(x)
}

func increment() {
	x++
}

/*
The two go routines are "main" and "increment".  The shared value is the global integer "x".
In main, the "increment" routine is invoked twice, then the "x" value is read and printed.  Since these
steps are not expected to occur in sequence, we don't know what the last step will print.  In my test,
the value printed is "0".
*/

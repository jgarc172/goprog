// Package motion defines functions of motion
// prompts the user for parameter values and prints the displacement
package main

import "fmt"

func main() {
	a, v0, s0 := readParams()
	fn := GenDisplaceFn(a, v0, s0)
	t := readTime()
	fmt.Println(fn(t))
}

// Displacement as a function of time
type Displacement func(float64) float64

// GenDisplaceFn returns a Displacement function
func GenDisplaceFn(a, v0, s0 float64) (fn Displacement) {
	// s = ½ a t2 + vot + so
	fn = func(t float64) float64 {
		return 0.5*a*t*t + v0*t + s0
	}
	return
}

// readParams prompts for acceleration
// initial velocity and displacement
func readParams() (a, v0, s0 float64) {
	fmt.Print("acceleration: ")
	fmt.Scanln(&a)
	fmt.Print("initial velocity: ")
	fmt.Scanln(&v0)
	fmt.Print("initial displacement: ")
	fmt.Scanln(&s0)
	return
}

// readTime prompts for a time value
func readTime() (t float64) {
	fmt.Print("time: ")
	fmt.Scanln(&t)
	return
}

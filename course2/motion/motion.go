// Package motion defines functions of motion
// and prints displacements for two values of time
package main

import "fmt"

func main() {
	fn := GenDisplaceFn(10, 2, 1)
	fmt.Println(fn(3))
	fmt.Println(fn(5))
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

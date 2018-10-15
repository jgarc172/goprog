// Package motion defines functions of motion
package main

func main() {

}

type time float64

// Displacement as a function of time
type Displacement func(time) float64

// GenDisplaceFn returns a
func GenDisplaceFn(a, v0, s0 float64) Displacement {

	fn := func(time) float64 {
		return 0
	}
	return fn
}

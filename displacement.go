package main

import (
	"fmt"
)

func GenDisplaceFn(acc, u0, s0 float64) func(t float64) float64 {
	f := func(t float64) float64 {
		return (0.5 * acc * t * t) + (u0 * t) + s0
	}
	return f
}

func main() {

	var acc, u0, s0, t float64
	fmt.Println("Enter the acceleration")
	fmt.Scanf("%f", &acc)
	fmt.Println("Enter the initial velocity")
	fmt.Scanf("%f", &u0)
	fmt.Println("Enter the initial displacement")
	fmt.Scanf("%f", &s0)
	fmt.Println("Enter the time")
	fmt.Scanf("%f", &t)
	fmt.Println(acc, u0, s0, t)
	fn := GenDisplaceFn(acc, u0, s0)
	fmt.Println(fn(t))
}

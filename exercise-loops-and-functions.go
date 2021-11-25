package main

import (
	"fmt"
)

func Sqrt(x float64) float64 {
	z := 1.0
	ep := 0.001
	times := 1
	for times < 1000 {
		z = z - (z*z-x)/(2*z)
		if z*z-x < ep {
			return z
		}
	}
	return z
}
func main() {
	fmt.Println(Sqrt(1))
}

package main

import (
	"fmt"
	"math"
)

//newton 's method
func Sqrt(x float64) float64 {
	z, d := float64(1), float64(1)
	for d > 1e-5 {
		y := z
		z -= (z*z - x) / (2 * z)
		d = math.Abs(z - y)
	}
	return z
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(math.Sqrt(2))
}

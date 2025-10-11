package main

import (
	"fmt"
)

func Sqrt(x float64) float64 {
	z := float64(1)

	for i := range 10 {
		z -= (z*z - x) / (2 * z)
		fmt.Printf("iter=%d z=%.12f\n", i+1, z)
	}

	return z
}

func main() {
	fmt.Println(Sqrt(2))
}

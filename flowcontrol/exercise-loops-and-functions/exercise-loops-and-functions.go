package main

import (
	"fmt"
	"math"
)

const e = 1e-8

func sqrt(x float64) float64 {
	z := 1.0

	for {
		fmt.Println(z)
		new_z := z - ((z*z - x) / (2 * z))
		if math.Abs(new_z-z) < e {
			return new_z
		}
		z = new_z
	}
}

func main() {
	fmt.Println(sqrt(2))
}

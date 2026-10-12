//practice exercise
package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := float64(1)
	prevZ := float64(0)
	for i := 0; i < 10; i++ {
		prevZ = z
		z -= (z*z - x) / (2*z)
		fmt.Println(z)
		if z == prevZ {
			break
		}
	}

	return z
}

func main() {
	realAns := math.Sqrt(2)
	fmt.Println("Correct answer:", realAns)
	fmt.Println(Sqrt(2))
}



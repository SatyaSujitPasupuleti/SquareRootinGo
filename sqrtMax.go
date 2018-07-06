package main

import (
	"fmt"
)

func Sqrt(x float64) float64 {
	z := 1.0
	count := 0
	for !(x-(z*z) == 0.0000) {
		z -= ((z*z - x) / (2 * z))
		count += 1

	}
	fmt.Println(count)
	return z
}
func main() {
	fmt.Println(Sqrt(9))

}

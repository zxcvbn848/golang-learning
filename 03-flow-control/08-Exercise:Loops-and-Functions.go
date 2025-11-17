// https://go.dev/tour/flowcontrol/8
package flowControl

import (
	"fmt"
)

func Sqrt(x float64) float64 {
	// Newton's method for finding square root
	// Start with an initial guess
	z := 1.0

	// Iterate until we get a good approximation
	for i := 0; i < 10; i++ {
		// Newton's method: z = z - (z*z - x) / (2*z)
		z = z - (z*z-x)/(2*z)
	}

	return z
}

func RunFlowControl08() {
	fmt.Println(Sqrt(7))
}

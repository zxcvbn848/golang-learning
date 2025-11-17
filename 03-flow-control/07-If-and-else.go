// https://go.dev/tour/flowcontrol/7
package flowControl

import (
	"fmt"
	"math"
)

func pow07(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	// can't use v here, though
	return lim
}

func RunFlowControl07() {
	fmt.Println(
		pow07(3, 2, 10),
		pow07(3, 3, 20),
	)
}

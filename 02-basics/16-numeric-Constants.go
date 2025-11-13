// https://go.dev/tour/basics/16
package basics

import (
	"fmt"
)

const (
	// Create a huge number by shifting a 1 bit left 100 places.
	// In other words, the binary number that is 1 followed by 100 zeroes.
	Big = 1 << 100
	// Shift it right again 99 places, so we end up with 1<<1, or 2.
	Small = Big >> 99
)

func needInt(x int) int { return x*10 + 1 }
func needFloat(x float64) float64 {
	return x * 0.1
}

func RunBasics16() {
	fmt.Println(needInt(Small))
	// ! 02-basics/16-numeric-Constants.go:23:22: cannot use Big (untyped int constant 1267650600228229401496703205376) as int value in argument to needInt (overflows)
	// fmt.Println(needInt(Big))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
}

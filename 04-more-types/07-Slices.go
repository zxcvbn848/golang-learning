// https://go.dev/tour/moretypes/7
package moreTypes

import (
	"fmt"
)

func RunMoreTypes07() {
	primes := [6]int{2, 3, 5, 7, 11, 13}

	var s []int = primes[1:4]
	fmt.Println(s) // [3 5 7]
}

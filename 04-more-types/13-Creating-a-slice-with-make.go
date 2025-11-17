// https://go.dev/tour/moretypes/13
package moreTypes

import (
	"fmt"
)

func RunMoreTypes13() {
	a := make([]int, 5)
	printSlice13("a", a) // a len=5 cap=5 [0 0 0 0 0]

	b := make([]int, 0, 5)
	printSlice13("b", b) // b len=0 cap=5 []

	c := b[:2]
	printSlice13("c", c) // c len=2 cap=5 [0 0]

	d := c[2:5]
	printSlice13("d", d) // d len=3 cap=3 [0 0 0]
}

func printSlice13(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}

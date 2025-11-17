// https://go.dev/tour/moretypes/8
package moreTypes

import (
	"fmt"
)

func RunMoreTypes08() {
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names) // [John Paul George Ringo]

	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b) // [John Paul] [Paul George]

	b[0] = "XXX"
	// slice 共享相同的底層 array，所以修改 b[0] 會影響 a[1] 和 names[1]
	fmt.Println(a, b)  // [John XXX] [XXX George]
	fmt.Println(names) // [John XXX George Ringo]
}

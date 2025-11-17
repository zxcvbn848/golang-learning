// https://go.dev/tour/moretypes/1
package moreTypes

import (
	"fmt"
)

func RunMoreTypes12() {
	var s []int
	fmt.Println(s, len(s), cap(s)) // [] 0 0
	// nil slice：長度為 0，容量為 0，沒有底層 array
	if s == nil {
		fmt.Println("nil!")
	} // nil!
}

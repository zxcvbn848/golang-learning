// https://go.dev/tour/moretypes/26
package moreTypes

import (
	"fmt"
)

func fibonacci() func() int {
	a, b := 0, 1
	return func() int {
		result := a
		a, b = b, a+b
		return result
	}
}

func RunMoreTypes26() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}

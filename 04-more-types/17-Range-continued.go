// https://go.dev/tour/moretypes/1
package moreTypes

import (
	"fmt"
)

func RunMoreTypes17() {
	pow := make([]int, 10)
	// 只使用 index
	for i := range pow {
		pow[i] = 1 << uint(i) // == 2**i
	}
	// 只使用 value
	for _, value := range pow {
		fmt.Printf("%d\n", value) // 1 2 4 8 16 32 64 128 256 512
	}
}

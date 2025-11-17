// https://go.dev/tour/moretypes/15
package moreTypes

import (
	"fmt"
)

func RunMoreTypes15() {
	var s []int
	printSlice15(s) // len=0 cap=0 []

	// append works on nil slices.
	s = append(s, 0)
	printSlice15(s) // len=1 cap=1 [0]

	// The slice grows as needed.
	s = append(s, 1)
	printSlice15(s) // len=2 cap=2 [0 1]

	// We can add more than one element at a time.
	// 當前的 len=2, cap=2，需要添加 3 個元素（總共需要 5 個位置）
	// Go 的 slice 擴容策略：當容量不足時會自動擴容
	// 從 cap=2 擴容時，Go 會分配 cap=6（而不是 cap=4 或 cap=5）
	// 這是為了預留空間，減少後續擴容的次數，提高效能
	s = append(s, 2, 3, 4)
	printSlice15(s) // len=5 cap=6 [0 1 2 3 4]
}

func printSlice15(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

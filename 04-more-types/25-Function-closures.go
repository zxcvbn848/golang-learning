// https://go.dev/tour/moretypes/25
package moreTypes

import (
	"fmt"
)

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func RunMoreTypes25() {
	/*
		這樣寫的好處
		- adder 回傳的函式可以獨立累加自己的 sum，pos 與 neg 各有自己的累積值，互不影響。
		- 使用者呼叫 pos(i) 或 neg(-2*i) 時，不需要自己管理 sum，closure 自動替你記住之前的結果。
		- 這種做法讓狀態被封裝在函式內部，API 更簡潔，也避免全域變數或指標等額外資料結構。
	*/
	pos, neg := adder(), adder()

	for i := range make([]int, 10) {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
		/*
			0 0
			1 -2
			3 -6
			6 -12
			10 -20
			15 -30
			21 -42
			28 -56
			36 -72
			45 -90
		*/
	}
}

// https://go.dev/tour/methods/13
// Nil interface values
// 重要概念：
// 1. nil interface value 既不持有值，也不持有具體型別（(nil, nil)）
// 2. 在 nil interface value 上呼叫方法會產生 runtime error（panic）
// 3. 因為沒有型別資訊，Go 無法知道要呼叫哪個具體方法
//
// 對比：
// - 前一個例子（12）：interface value 是 (nil, *T12)，有型別資訊，可以呼叫方法
// - 這個例子：interface value 是 (nil, nil)，沒有型別資訊，無法呼叫方法
package methods

import (
	"fmt"
)

type I13 interface {
	M()
}

func RunMethods13() {
	// 宣告一個 interface variable，但沒有賦值
	// 此時 i 是 nil interface value：(nil, nil)
	var i I13
	Describe13(i) // 輸出：(<nil>, <nil>) - 既沒有值，也沒有型別

	// 在 nil interface value 上呼叫方法會產生 runtime error (panic)
	// 因為 Go 不知道要呼叫哪個具體型別的方法
	// 錯誤訊息：panic: runtime error: invalid memory address or nil pointer dereference
	// 注意：linter 會警告這是 nil dereference，但這是教學範例，用來展示這個錯誤
	i.M() // 這行會導致 panic
}

// Describe13 展示 interface value 的內部結構
func Describe13(i I13) {
	fmt.Printf("(%v, %T)\n", i, i)
}

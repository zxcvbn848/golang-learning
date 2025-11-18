// https://go.dev/tour/methods/12
// Interface values with nil underlying values
// 重要概念：
// 1. 如果 interface value 中的具體值是 nil，方法仍會被呼叫（但 receiver 是 nil）
// 2. 在 Go 中，可以寫出優雅處理 nil receiver 的方法（不會像其他語言拋出 null pointer exception）
// 3. 注意：即使 interface value 中持有的具體值是 nil，interface value 本身並不是 nil
package methods

import (
	"fmt"
)

type I12 interface {
	M()
}

type T12 struct {
	S string
}

// M 方法優雅地處理 nil receiver 的情況
// 在 Go 中，即使 receiver 是 nil，方法仍會被呼叫
func (t *T12) M() {
	if t == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Println(t.S)
}

func RunMethods12() {
	var i I12

	// 創建一個 nil 指標
	var t *T12
	// 將 nil 指標賦值給 interface value
	// 此時 i 不是 nil，而是 (nil, *methods.T12)
	i = t
	Describe12(i) // 輸出：(<nil>, *methods.T12) - interface value 本身不是 nil！
	i.M()         // 可以安全呼叫，因為 M 方法處理了 nil 的情況

	// 對比：正常的指標
	i = &T12{"hello"}
	Describe12(i) // 輸出：(&{hello}, *methods.T12)
	i.M()         // 輸出：hello
}

// Describe12 展示 interface value 的內部結構
func Describe12(i I12) {
	fmt.Printf("(%v, %T)\n", i, i)
}

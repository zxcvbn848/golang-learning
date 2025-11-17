// https://go.dev/tour/methods/10
package methods

import (
	"fmt"
)

// 定義一個介面，只需列出方法簽章
type I interface {
	M()
}

type T struct {
	S string
}

// 只要型別 T 實作了介面所需的方法（M），就被視為實作 I
// 不需要像其他語言一樣用 implements/extends 關鍵字
func (t T) M() {
	fmt.Println(t.S)
}

func RunMethods10() {
	// 將 T 指派給介面，不需額外宣告「T 實作 I」
	var i I = T{"hello"}
	i.M() // hello
}

// https://go.dev/tour/methods/14
// Empty interface（interface{}）
// - 定義：沒有任何方法的介面
// - 特性：所有型別都至少實作 0 個方法，因此任何型別都符合空介面
// - 用途：當需要接受「任意型別」時，例如 fmt.Print 的參數
package methods

import (
	"fmt"
)

func RunMethods14() {
	var i interface{} // 空介面，可承載任何值
	Describe14(i)     // (<nil>, <nil>)

	i = 42
	Describe14(i) // (42, int)

	i = "hello"
	Describe14(i) // (hello, string)
}

// Describe14 顯示空介面目前包裝的值與型別
func Describe14(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}

// https://go.dev/tour/methods/15
// Type assertions
// - 語法：t := i.(T)  斷言 i 目前持有的具體型別是 T
// - 若斷言成功，回傳底層的 T 值；否則觸發 panic
// - 也可以使用兩個回傳值：t, ok := i.(T)，失敗時 ok=false、t 為 T 的零值
// - 語法與從 map 讀取值時的「value, ok := m[key]」類似
package methods

import (
	"fmt"
)

func RunMethods15() {
	var i interface{} = "hello"

	// 1. 直接斷言（單一回傳值），若失敗會 panic
	s := i.(string)
	fmt.Println(s) // hello

	// 2. 斷言並接收 ok，避免 panic
	s, ok := i.(string)
	fmt.Println(s, ok) // hello true

	f, ok := i.(float64)
	fmt.Println(f, ok) // 0 false（zero value + 失敗）

	// 3. 直接斷言錯誤型別會 panic
	f = i.(float64) // panic: interface conversion: interface {} is string, not float64
	fmt.Println(f)
}

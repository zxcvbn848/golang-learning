// https://go.dev/tour/methods/16
// Type switches
// Type switch 允許進行一系列型別斷言
// 語法類似一般的 switch，但 case 指定的是型別（不是值）
// 這些型別會與 interface value 中持有的值的型別進行比較
//
// 語法：switch v := i.(type) { ... }
// - 與 type assertion i.(T) 語法相似，但 T 被替換為關鍵字 type
// - 在每個 case 中，v 會是該 case 指定的型別
// - 在 default case 中，v 與 i 有相同的 interface 型別和值
package methods

import (
	"fmt"
)

func do(i interface{}) {
	// type switch：根據 i 的實際型別來執行不同的邏輯
	// v := i.(type) 會根據 i 的實際型別來賦值給 v
	switch v := i.(type) {
	case int:
		// 在這個 case 中，v 的型別是 int
		// 可以直接使用 int 的操作（如乘法）
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		// 在這個 case 中，v 的型別是 string
		// 可以使用 string 的操作（如 len()）
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		// 沒有匹配的型別時，v 與 i 有相同的 interface 型別
		// 可以使用 %T 來顯示型別資訊
		fmt.Printf("I don't know about type %T!\n", v)
	}
}

func RunMethods16() {
	do(21)      // int 型別，輸出：Twice 21 is 42
	do("hello") // string 型別，輸出："hello" is 5 bytes long
	do(true)    // bool 型別，輸出：I don't know about type bool!
}

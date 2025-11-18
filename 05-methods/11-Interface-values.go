// https://go.dev/tour/methods/11
// Interface values 的內部結構
// 在底層，interface value 可以視為一個 (value, type) 的元組：
// - value: 具體的值（concrete value）
// - type: 具體的型別（concrete type）
// 當我們在 interface value 上呼叫方法時，Go 會根據其底層的具體型別來執行對應的方法
package methods

import (
	"fmt"
	"math"
)

type I11 interface {
	M()
}

type T11 struct {
	S string
}

func (t *T11) M() {
	fmt.Println(t.S)
}

type F11 float64

func (f F11) M() {
	fmt.Println(f)
}

func RunMethods11() {
	var i I11

	// interface value 可以視為 (value, type) 的元組
	// 第一個值：具體的值（concrete value）
	// 第二個值：具體的型別（concrete type）
	i = &T11{"Hello"}
	Describe11(i) // 輸出：(&{Hello}, *methods.T11) - 值是指標，型別是 *T11
	i.M()         // 呼叫 *T11 的 M 方法

	i = F11(math.Pi)
	Describe11(i) // 輸出：(3.141592653589793, methods.F11) - 值是 float64，型別是 F11
	i.M()         // 呼叫 F11 的 M 方法
}

// Describe11 展示 interface value 的內部結構：(value, type)
func Describe11(i I11) {
	fmt.Printf("(%v, %T)\n", i, i)
}

// https://go.dev/tour/methods/3
package methods

import (
	"fmt"
	"math"
)

// MyFloat 是自訂的數值型別（非 struct），仍然可以擁有方法
// 注意：只有在同一個 package 定義的型別，才能為它宣告方法
type MyFloat float64

// Abs 方法示範對非 struct 型別宣告方法
func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func RunMethods03() {
	f := MyFloat(-math.Sqrt2)
	// 呼叫 MyFloat 的 Abs 方法，與 struct receiver 用法相同
	fmt.Println(f.Abs()) // 1.4142135623730951
}

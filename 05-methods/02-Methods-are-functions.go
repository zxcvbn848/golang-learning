// https://go.dev/tour/methods/2
package methods

import (
	"fmt"
	"math"
)

type Vertex02 struct {
	X, Y float64
}

// Abs 這裡寫成一般函數（以 Vertex02 當參數），與方法版本功能相同
// 說明「方法其實就是帶 receiver 的函數」，只是這裡把 receiver 改成顯式參數
func Abs(v Vertex02) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func RunMethods02() {
	v := Vertex02{3, 4}
	// 直接呼叫一般函數 Abs，計算結果與方法版本完全一致
	fmt.Println(Abs(v)) // 5
}

// https://go.dev/tour/methods/7
package methods

import (
	"fmt"
	"math"
)

type Vertex07 struct {
	X, Y float64
}

// Abs 使用值接收者，因此既可以由值呼叫，也可以由指標（會被自動解參考）呼叫
func (v Vertex07) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// AbsFunc 則是單純的函數，只接受 Vertex07 值本身
// 嘗試呼叫 AbsFunc(&v) 會編譯錯誤，因為它期望的是值不是指標
func AbsFunc(v Vertex07) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func RunMethods07() {
	v := Vertex07{3, 4}
	fmt.Println(v.Abs())    // 5
	fmt.Println(AbsFunc(v)) // 5

	p := &Vertex07{4, 3}
	// 指標也能呼叫值 receiver 方法，Go 會自動進行 (*p).Abs()
	fmt.Println(p.Abs()) // 5
	// 但函數仍需要傳值，因此必須顯式寫成 AbsFunc(*p)
	fmt.Println(AbsFunc(*p)) // 5
}

// https://go.dev/tour/methods/9
package methods

import (
	"fmt"
	"math"
)

type Abser09 interface {
	Abs() float64
}

func RunMethods09() {
	var a Abser09
	f := MyFloat09(-math.Sqrt2)
	v := Vertex09{3, 4}

	a = f                // MyFloat09 有 Abs 方法（值 receiver），可直接賦值
	fmt.Println(a.Abs()) // 1.4142135623730951
	a = &v               // 只有 *Vertex09 實作 Abs，因此要給指標

	// In the following line, v is a Vertex (not *Vertex)
	// and does NOT implement Abser.
	// 取消註解會編譯錯誤： Vertex09 缺少 Abs 方法
	// a = v

	fmt.Println(a.Abs()) // 5
}

type MyFloat09 float64

func (f MyFloat09) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

type Vertex09 struct {
	X, Y float64
}

func (v *Vertex09) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

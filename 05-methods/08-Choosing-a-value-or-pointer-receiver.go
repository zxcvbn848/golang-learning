// https://go.dev/tour/methods/8
package methods

import (
	"fmt"
	"math"
)

type Vertex08 struct {
	X, Y float64
}

// 選擇指標 receiver 的兩大理由：
// 1. 方法需要修改原始資料（Scale 會調整座標值）
// 2. 避免在每次呼叫時複製大型結構，提高效率
func (v *Vertex08) Scale(f float64) {
	v.X *= f
	v.Y *= f
}

// 即便 Abs 只讀資料，仍使用指標 receiver，好處是與其他方法保持一致，避免混用值/指標 receiver
func (v *Vertex08) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func RunMethods08() {
	v := &Vertex08{3, 4}
	fmt.Printf("Before scaling: %+v, Abs: %v\n", v, v.Abs())
	v.Scale(5)
	fmt.Printf("After scaling: %+v, Abs: %v\n", v, v.Abs())
}

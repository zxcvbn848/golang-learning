// https://go.dev/tour/methods/4
package methods

import (
	"fmt"
	"math"
)

type Vertex04 struct {
	X, Y float64
}

// Abs 仍使用值接收者，因為它只需讀取數值
func (v Vertex04) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// Scale 使用指標接收者 (*Vertex04)，才能直接修改呼叫者的內容
// 若將這裡的 * 移除，Scale 只會修改副本，原始 v 不會被改變
func (v *Vertex04) Scale(f float64) {
	v.X *= f
	v.Y *= f
}

func RunMethods04() {
	v := Vertex04{3, 4}
	// 指標接收者允許我們在方法內直接調整 v 的值
	v.Scale(10)
	fmt.Println(v.Abs()) // 50
}

// https://go.dev/tour/methods/1
package methods

import (
	"fmt"
	"math"
)

// Vertex 是一個自訂型別（struct），我們可以在它上面定義方法
type Vertex struct {
	X, Y float64
}

// Abs 是定義在 Vertex 型別上的方法（receiver 是 v Vertex）
// receiver 放在 func 與方法名稱之間，讓這個函式綁定到 Vertex 型別
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func RunMethods01() {
	// 雖然 Go 沒有 class，但可以透過 receiver 讓 Vertex 擁有方法
	v := Vertex{3, 4}
	fmt.Println(v.Abs()) // 5
}

// https://go.dev/tour/methods/5
package methods

import (
	"fmt"
	"math"
)

type Vertex05 struct {
	X, Y float64
}

// Abs05 與方法版本相同，只是寫成一般函數
func Abs05(v Vertex05) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// Scale05 以 *Vertex05 為參數，模擬「指標 receiver」的行為
// 若把這裡的 * 拿掉，就只能修改副本，呼叫端仍需傳入值類型並無法改變原來的 Vertex05
// 同時，呼叫端就必須改成傳值（而不是 &v）才能編譯，這展示為何指標參數更適合用來修改原值
func Scale05(v *Vertex05, f float64) {
	v.X *= f
	v.Y *= f
}

func RunMethods05() {
	v := Vertex05{3, 4}
	// 傳入 &v 讓 Scale05 能直接修改 v 的內容
	Scale05(&v, 10)
	fmt.Println(Abs05(v)) // 50
}

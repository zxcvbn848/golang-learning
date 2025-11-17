// https://go.dev/tour/methods/6
package methods

import (
	"fmt"
)

type Vertex06 struct {
	X, Y float64
}

// Scale 是指標接收者方法，可透過值或指標來呼叫
// v.Scale(5) 會被 Go 自動轉成 (&v).Scale(5)
func (v *Vertex06) Scale(f float64) {
	v.X *= f
	v.Y *= f
}

// ScaleFunc 是對應的函數版本：要手動傳入 *Vertex06 才能修改原值
// 如果誤用 ScaleFunc(v, 5)（傳值），會無法編譯
func ScaleFunc(v *Vertex06, f float64) {
	v.X *= f
	v.Y *= f
}

func RunMethods06() {
	v := Vertex06{3, 4}
	// 雖然 v 是值，但仍可呼叫指標 receiver 方法（Go 會自動取址）
	v.Scale(2)
	// 函數版本必須手動傳入指標，否則編譯錯誤
	ScaleFunc(&v, 10)

	p := &Vertex06{4, 3}
	// 直接用指標呼叫沒問題
	p.Scale(3)
	ScaleFunc(p, 8)

	fmt.Println(v, p) // {60 80} &{96 72}
}

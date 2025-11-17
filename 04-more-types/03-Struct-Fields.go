// https://go.dev/tour/moretypes/3
package moreTypes

import (
	"fmt"
)

type Vertex03 struct {
	X int
	Y int
}

func RunMoreTypes03() {
	v := Vertex03{1, 2}
	// 寫入新的值到 X
	v.X = 4
	// 取得新的 X 值
	fmt.Println(v.X) // 4
}

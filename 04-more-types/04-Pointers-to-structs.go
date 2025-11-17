// https://go.dev/tour/moretypes/4
package moreTypes

import (
	"fmt"
)

type Vertex04 struct {
	X int
	Y int
}

func RunMoreTypes04() {
	v := Vertex04{1, 2}
	p := &v
	// 理論上應寫 *p.X，但可省略成 p.X
	p.X = 1e9
	fmt.Println(v) // {1000000000 2}
}

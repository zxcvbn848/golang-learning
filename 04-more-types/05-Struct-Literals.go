// https://go.dev/tour/moretypes/5
package moreTypes

import (
	"fmt"
)

type Vertex05 struct {
	X, Y int
}

var (
	v1 = Vertex05{1, 2}  // has type Vertex
	v2 = Vertex05{X: 1}  // Y:0 is implicit
	v3 = Vertex05{}      // X:0 and Y:0
	p  = &Vertex05{1, 2} // has type *Vertex
)

func RunMoreTypes05() {
	fmt.Println(v1, p, v2, v3) // {1 2} &{1 2} {1 0} {0 0}
}

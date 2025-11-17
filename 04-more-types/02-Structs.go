// https://go.dev/tour/moretypes/2
package moreTypes

import (
	"fmt"
)

type Vertex struct {
	X int
	Y int
}

func RunMoreTypes02() {
	fmt.Println(Vertex{1, 2})
}

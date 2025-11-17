// https://go.dev/tour/moretypes/1
package moreTypes

import (
	"fmt"
)

type Vertex21 struct {
	Lat, Long float64
}

var m21 = map[string]Vertex21{
	// 可以省略 Vertex21{}
	"Bell Labs": {40.68433, -74.39967},
	"Google":    {37.42202, -122.08408},
}

func RunMoreTypes21() {
	fmt.Println(m21)
}

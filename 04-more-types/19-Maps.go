// https://go.dev/tour/moretypes/19
package moreTypes

import (
	"fmt"
)

type Vertex19 struct {
	Lat, Long float64
}

var m map[string]Vertex19

func RunMoreTypes19() {
	m = make(map[string]Vertex19)
	m["Bell Labs"] = Vertex19{
		40.68433, -74.39967,
	}
	fmt.Println(m["Bell Labs"]) // {40.68433 -74.39967}
}

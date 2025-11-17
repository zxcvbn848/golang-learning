// https://go.dev/tour/moretypes/22
package moreTypes

import (
	"fmt"
)

func RunMoreTypes22() {
	m22 := make(map[string]int)

	m22["Answer"] = 42
	fmt.Println("The value:", m22["Answer"]) // The value: 42

	m22["Answer"] = 48
	fmt.Println("The value:", m22["Answer"]) // The value: 48

	delete(m22, "Answer")
	fmt.Println("The value:", m22["Answer"]) // The value: 0

	v, ok := m22["Answer"]
	fmt.Println("The value:", v, "Present?", ok) // The value: 0 Present? false
}

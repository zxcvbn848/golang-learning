// https://go.dev/tour/flowcontrol/5
package flowControl

import (
	"fmt"
	"math"
)

func sqrt(x float64) string {
	// 不需要小括號 () ，但大括號 {} 是必須的
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func RunFlowControl05() {
	fmt.Println(sqrt(2), sqrt(-4))
}

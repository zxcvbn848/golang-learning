// https://go.dev/tour/flowcontrol/6
package flowControl

import (
	"fmt"
	"math"
)

func pow(x, n, lim float64) float64 {
	// 在 if 語句中宣告一個變數，並在 if 語句中使用
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}

func RunFlowControl06() {
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)
}

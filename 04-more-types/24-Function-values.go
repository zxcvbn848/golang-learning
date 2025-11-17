// https://go.dev/tour/moretypes/24
package moreTypes

import (
	"fmt"
	"math"
)

func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

func RunMoreTypes24() {
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypot(5, 12)) // 13

	// compute 呼叫 hypot(3, 4) 計算 sqrt(3^2 + 4^2)
	fmt.Println(compute(hypot)) // 5
	// compute 呼叫 math.Pow(3, 4) 計算 3 的 4 次方
	fmt.Println(compute(math.Pow)) // 81
}

// https://go.dev/tour/basics/12
package basics

import (
	"fmt"
	"math"
)

func RunBasics13() {
	var x, y int = 3, 4
	// 如果移除了 float64() 會出錯，也就是不能只寫 math.Sqrt(x*x + y*y)
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)
	fmt.Println(x, y, z)
}

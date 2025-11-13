// https://go.dev/tour/flowcontrol/1
package flowControl

import (
	"fmt"
)

func RunFlowControl01() {
	sum := 0
	// 不需要小括號 ()，但大括號 {} 是必須的
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)
}

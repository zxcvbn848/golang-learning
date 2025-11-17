// https://go.dev/tour/flowcontrol/13
package flowControl

import (
	"fmt"
)

func RunFlowControl13() {
	// defer 會按照後進先出 (LIFO) 的順序執行
	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")
	/*
		輸出：
		counting
		done
		9
		8
		7
		6
		5
		4
		3
		2
		1
		0
	*/
}

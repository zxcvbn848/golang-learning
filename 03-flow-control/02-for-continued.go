// https://go.dev/tour/flowcontrol/2
package flowControl

import (
	"fmt"
)

func RunFlowControl02() {
	sum := 1
	// 可以省略 i:=0; & i++ & 分號
	// 分號可省略是 03-for-is-Go's-"while".go 的內容，但因為 go 儲存時會自動移除分號，所以放在一起
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}

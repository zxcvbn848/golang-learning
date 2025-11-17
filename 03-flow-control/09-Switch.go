// https://go.dev/tour/flowcontrol/9
package flowControl

import (
	"fmt"
	"runtime"
)

func RunFlowControl09() {
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	/*
		case 的值不需要是常數，也不需要是 integer
		- 字串
		- 布林值
		- rune / byte
		- 自訂型別
		- 比較運算結果
		- interface
	*/
	case "darwin":
		fmt.Println("macOS.")
		// 內建 break
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.\n", os)
	}
}

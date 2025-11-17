// https://go.dev/tour/flowcontrol/12
package flowControl

import (
	"fmt"
)

/*
defer 常見用途
- 關閉檔案
- 解鎖 mutex
- 記錄執行時間
- panic 時仍能收尾
*/
func RunFlowControl12() {
	// defer 語句會將函數的執行順序延後到外層函數返回之前
	defer fmt.Println("world")

	fmt.Println("hello")
	// 輸出: hello world
	// --------------------------------
	// defer 會「捕捉」當下的參數值
	x := 10
	defer fmt.Println(x)
	x = 20
	// 輸出: 10
	// --------------------------------
	// 如果想捕捉變動後的值，需要用 closure：
	y := 10
	defer func() { fmt.Println(y) }()
	y = 20
	// 輸出: 20
}

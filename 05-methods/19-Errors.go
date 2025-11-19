// https://go.dev/tour/methods/19
package methods

import (
	"fmt"
	"time"
)

// MyError 實作內建的 error 介面（只有一個 Error() string 方法）。
// Go 透過 error 值表示失敗狀態，nil 代表成功、非 nil 代表失敗。
type MyError struct {
	When time.Time
	What string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s",
		e.When, e.What)
}

func run() error {
	// 實際情況會在發生錯誤時回傳具體的錯誤內容。
	// 此處示範建立 MyError 並回傳，呼叫端可檢查是否為 nil。
	return &MyError{
		time.Now(),
		"it didn't work",
	}
}

func RunMethods19() {
	// 典型錯誤處理流程：呼叫函式→檢查 error 是否為 nil。
	if err := run(); err != nil {
		// fmt 會自動呼叫 err.Error()，顯示人類可讀的訊息。
		fmt.Println(err)
	}
}

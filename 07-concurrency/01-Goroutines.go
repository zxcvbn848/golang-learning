// https://go.dev/tour/concurrency/1
package concurrency

import (
	"fmt"
	"time"
)

// Goroutine（協程）是由 Go 運行時管理的輕量級線程
// 與傳統操作系統線程相比，goroutine 的創建和切換開銷非常小
// 一個程序可以輕鬆創建成千上萬個 goroutine

// say 函數會打印字符串 s 五次，每次打印之間暫停 100 毫秒
func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func RunConcurrency01() {
	// 使用 go 關鍵字啟動一個新的 goroutine
	// 語法：go f(x, y, z)
	//
	// 重要說明：
	// 1. f, x, y, z 的評估（evaluation）發生在當前的 goroutine 中
	//    也就是說，參數值在啟動新 goroutine 之前就已經計算好了
	// 2. 函數 f 的執行（execution）發生在新的 goroutine 中
	//    新 goroutine 會異步執行，不會阻塞當前 goroutine
	//
	// 在這個例子中：
	// - "world" 參數在當前 goroutine 中評估
	// - say("world") 的執行在新的 goroutine 中進行
	//
	// 主要區別：
	// ┌─────────────────────────────────────────────────────────┐
	// │ go say("world")  vs  say("hello")                       │
	// ├─────────────────────────────────────────────────────────┤
	// │                                                          │
	// │ go say("world"):                                         │
	// │   ✓ 非阻塞：立即返回，不等待函數執行完成                │
	// │   ✓ 異步執行：在新的 goroutine 中運行                   │
	// │   ✓ 並發：與當前 goroutine 同時執行                     │
	// │                                                          │
	// │ say("hello"):                                            │
	// │   ✓ 阻塞：等待函數執行完成後才繼續                      │
	// │   ✓ 同步執行：在當前 goroutine 中運行                   │
	// │   ✓ 順序：必須等這行執行完才能執行後續代碼             │
	// │                                                          │
	// └─────────────────────────────────────────────────────────┘
	go say("world") // 啟動新 goroutine，非阻塞，立即繼續執行下一行

	// 這行代碼在當前的 goroutine 中執行（主 goroutine）
	// 此時兩個 goroutine 會並發運行：
	//   - 主 goroutine 執行 say("hello")
	//   - 新 goroutine 執行 say("world")
	// 輸出順序可能交錯，例如：hello, world, hello, world, ...
	say("hello") // 同步執行，阻塞直到完成

	// 注意：如果主 goroutine 結束，所有其他 goroutine 也會被終止
	// 為了確保新啟動的 goroutine 有時間執行，我們需要等待
	// 在這個例子中，say("hello") 會執行 5 次，每次 100ms，總共約 500ms
	// 這給了 say("world") goroutine 足夠的時間完成執行

	// 更好的做法是使用 sync.WaitGroup 或 channel 來同步
	// 但這裡為了簡單演示，依賴時間來確保 goroutine 完成
	time.Sleep(600 * time.Millisecond)

	// Goroutines 的重要特性：
	// 1. 輕量級：創建和切換開銷很小，可以創建大量 goroutine
	// 2. 共享地址空間：所有 goroutine 運行在同一個進程的地址空間中
	// 3. 需要同步：由於共享內存，訪問共享數據時必須進行同步
	// 4. sync 包提供了同步原語（如 Mutex, WaitGroup 等）
	//    但在 Go 中，通常使用 channel 進行通信會更好（見後續章節）
}

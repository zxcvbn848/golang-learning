// https://go.dev/tour/concurrency/5
package concurrency

import (
	"fmt"
	"time"
)

// Select 語句
//
// select 語句讓一個 goroutine 可以等待多個通信操作
//
// 工作原理：
// - select 會阻塞，直到其中一個 case 可以運行
// - 如果多個 case 都準備好了，它會隨機選擇一個執行
// - 這使得 goroutine 可以同時監聽多個 channel
//
// 語法：
//   select {
//   	case <-ch1:
// 			處理 ch1 的接收
//   	case ch2 <- value:
// 			處理 ch2 的發送
//   	case <-time.After(time.Second):
// 			超時處理
//   	default:
// 			如果所有 case 都阻塞，執行 default（可選）
//   }

// fibonacci05 函數生成斐波那契數列
// 使用 select 語句同時處理發送和接收操作
func fibonacci05(c, quit chan int) {
	x, y := 0, 1
	for {
		// select 語句會等待以下兩個 case 中的一個可以執行：
		select {
		case c <- x:
			// 如果可以向 channel c 發送值（接收者準備好了）
			// 發送當前的斐波那契數，然後計算下一個
			x, y = y, x+y
		case <-quit:
			// 如果從 quit channel 接收到值（收到退出信號）
			// 打印退出消息並返回，結束函數
			fmt.Println("quit")
			return
		}
		// 注意：select 會阻塞，直到其中一個 case 可以執行
		// 如果兩個 case 都準備好了，會隨機選擇一個
	}
}

func RunConcurrency05() {
	fmt.Println("=== 示例 1: 使用 Select 處理多個 Channel ===")

	c := make(chan int)
	quit := make(chan int)

	// 啟動一個 goroutine 來接收斐波那契數
	go func() {
		// 接收 10 個斐波那契數
		for range 10 {
			fmt.Println(<-c)
		}
		// 接收完 10 個數後，發送退出信號
		quit <- 0
	}()

	// 在主 goroutine 中運行 fibonacci05
	// 它會持續生成斐波那契數，直到收到退出信號
	fibonacci05(c, quit)

	fmt.Println()

	fmt.Println("=== 示例 2: Select 的隨機選擇行為 ===")

	ch1 := make(chan string)
	ch2 := make(chan string)

	// 啟動兩個 goroutine 同時發送數據
	go func() {
		ch1 <- "來自 ch1"
	}()
	go func() {
		ch2 <- "來自 ch2"
	}()

	// 如果兩個 channel 都準備好了，select 會隨機選擇一個
	// 多次運行可能會看到不同的順序
	select {
	case msg1 := <-ch1:
		fmt.Println("接收到:", msg1)
	case msg2 := <-ch2:
		fmt.Println("接收到:", msg2)
	}

	// 接收另一個值（如果還有）
	select {
	case msg1 := <-ch1:
		fmt.Println("接收到:", msg1)
	case msg2 := <-ch2:
		fmt.Println("接收到:", msg2)
	}

	fmt.Println()

	fmt.Println("=== 示例 3: Select 與 Default（非阻塞） ===")

	ch := make(chan int)

	// 使用 default case 可以實現非阻塞的發送或接收
	select {
	case ch <- 1:
		fmt.Println("發送成功")
	case <-ch:
		fmt.Println("接收成功")
	default:
		// 如果所有 case 都阻塞，立即執行 default
		// 這使得 select 不會阻塞
		fmt.Println("Channel 未準備好，執行 default")
	}

	fmt.Println()

	fmt.Println("=== 示例 4: Select 與超時 ===")

	ch3 := make(chan string)

	// 啟動一個 goroutine，延遲發送
	go func() {
		time.Sleep(2 * time.Second)
		ch3 <- "數據"
	}()

	// 使用 time.After 實現超時機制
	select {
	case msg := <-ch3:
		fmt.Println("接收到:", msg)
	case <-time.After(1 * time.Second):
		// 1 秒後如果還沒收到數據，執行超時處理
		fmt.Println("超時：1 秒內未收到數據")
	}

	fmt.Println()

	fmt.Println("=== 總結 ===")
	fmt.Println("1. select 允許等待多個 channel 操作")
	fmt.Println("2. select 會阻塞直到一個 case 可以執行")
	fmt.Println("3. 如果多個 case 都準備好，會隨機選擇一個")
	fmt.Println("4. 使用 default 可以實現非阻塞操作")
	fmt.Println("5. 常用於實現超時、取消和優先級處理")
}

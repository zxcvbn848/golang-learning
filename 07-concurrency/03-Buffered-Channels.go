// https://go.dev/tour/concurrency/3
package concurrency

import (
	"fmt"
)

// 緩衝通道（Buffered Channels）
//
// Channel 可以是帶緩衝的。提供緩衝區長度作為 make 的第二個參數來初始化緩衝通道：
//   ch := make(chan int, 100)
//
// 緩衝通道與無緩衝通道的區別：
// ┌─────────────────────────────────────────────────────────────┐
// │ 無緩衝通道 (make(chan int)):                                │
// │   - 發送操作會阻塞，直到有接收者準備好接收                  │
// │   - 接收操作會阻塞，直到有發送者準備好發送                  │
// │   - 要求發送和接收同時準備好（同步）                        │
// │                                                              │
// │ 緩衝通道 (make(chan int, n)):                               │
// │   - 發送到緩衝通道只有在緩衝區滿時才會阻塞                  │
// │   - 從緩衝通道接收只有在緩衝區空時才會阻塞                  │
// │   - 允許異步操作：發送者可以發送多個值而不等待接收者       │
// └─────────────────────────────────────────────────────────────┘

func RunConcurrency03() {
	fmt.Println("=== 示例 1: 正常使用緩衝通道 ===")

	// 創建一個緩衝區大小為 2 的通道
	// 這意味著可以發送 2 個值而不阻塞（只要緩衝區未滿）
	ch := make(chan int, 2)

	// 發送兩個值到緩衝通道
	// 因為緩衝區大小為 2，這兩個發送操作都不會阻塞
	ch <- 1
	ch <- 2
	fmt.Println("已發送 2 個值到緩衝通道（緩衝區大小為 2）")

	// 接收並打印值
	fmt.Println("接收值:", <-ch) // 1
	fmt.Println("接收值:", <-ch) // 2

	fmt.Println()

	fmt.Println("=== 示例 2: 緩衝區溢出（會導致死鎖） ===")

	// 創建一個緩衝區大小為 2 的通道
	ch2 := make(chan int, 2)

	// 發送兩個值（緩衝區滿了）
	ch2 <- 1
	ch2 <- 2
	fmt.Println("已發送 2 個值，緩衝區已滿")

	// 嘗試發送第三個值
	// 這會導致阻塞，因為緩衝區已滿且沒有接收者
	// 如果沒有其他 goroutine 來接收，程序會死鎖（deadlock）
	fmt.Println("嘗試發送第三個值（會阻塞，因為緩衝區已滿）...")

	// 注意：在實際運行時，這行會導致死鎖
	// 因為主 goroutine 會永遠阻塞在這裡，等待緩衝區有空間
	// 但沒有其他 goroutine 來接收值
	// ch2 <- 3  // 取消註釋這行會導致程序死鎖

	// 為了演示，我們先接收一個值，然後再發送
	fmt.Println("先接收一個值:", <-ch2) // 接收 1，緩衝區現在有空間
	ch2 <- 3                      // 現在可以發送第三個值了
	fmt.Println("成功發送第三個值")

	// 接收剩餘的值
	fmt.Println("接收值:", <-ch2) // 2
	fmt.Println("接收值:", <-ch2) // 3

	fmt.Println()

	fmt.Println("=== 示例 3: 緩衝區為空時接收會阻塞 ===")

	ch3 := make(chan int, 2)

	// 緩衝區為空時嘗試接收會阻塞
	// 如果沒有其他 goroutine 來發送，程序會死鎖
	fmt.Println("緩衝區為空，嘗試接收會阻塞...")

	// 為了演示，我們先發送一個值
	ch3 <- 10
	fmt.Println("已發送值 10")

	// 現在可以接收了
	fmt.Println("接收值:", <-ch3) // 10

	// 如果緩衝區再次為空，接收會阻塞
	// fmt.Println(<-ch3)  // 取消註釋這行會導致死鎖

	fmt.Println()

	fmt.Println("=== 總結 ===")
	fmt.Println("緩衝通道的行為：")
	fmt.Println("  - 發送：只有在緩衝區滿時才阻塞")
	fmt.Println("  - 接收：只有在緩衝區空時才阻塞")
	fmt.Println("  - 死鎖：當所有 goroutine 都在等待時發生（發送者等待空間，接收者等待數據）")

	// RunConcurrency03Deadlock()
}

// RunConcurrency03Deadlock 演示緩衝區溢出導致的死鎖
// 警告：這個函數會導致程序死鎖，僅用於演示目的
// 要運行此函數，請取消註釋 RunConcurrency03Deadlock() 的調用
func RunConcurrency03Deadlock() {
	fmt.Println("=== 演示：緩衝區溢出導致死鎖 ===")

	// 創建緩衝區大小為 2 的通道
	ch := make(chan int, 2)

	// 發送兩個值，緩衝區已滿
	ch <- 1
	ch <- 2
	fmt.Println("已發送 2 個值，緩衝區已滿（大小為 2）")

	// 嘗試發送第三個值
	// 這會導致主 goroutine 永遠阻塞在這裡
	// 因為緩衝區已滿，且沒有其他 goroutine 來接收值
	// 程序會死鎖，Go 運行時會檢測到並報錯：
	// "fatal error: all goroutines are asleep - deadlock!"
	fmt.Println("嘗試發送第三個值（會導致死鎖）...")
	ch <- 3 // 這行會導致死鎖

	// 這行永遠不會執行
	fmt.Println("這行永遠不會執行")
}

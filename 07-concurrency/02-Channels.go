// https://go.dev/tour/concurrency/2
package concurrency

import (
	"fmt"
)

// Channel（通道）是一個類型化的管道，用於在 goroutine 之間傳遞數據
// 使用通道運算符 <- 來發送和接收值
//
// 語法說明：
//   ch <- v    // 將 v 發送到 channel ch（數據流向箭頭方向）
//   v := <-ch  // 從 ch 接收值，並賦值給 v
//
// 與 map 和 slice 一樣，channel 必須在使用前創建：
//   ch := make(chan int)
//
// 重要特性：
// - 默認情況下，發送和接收操作會阻塞，直到另一端準備好
//   這使得 goroutine 可以在沒有顯式鎖或條件變量的情況下同步
// - Channel 是類型安全的，只能傳遞指定類型的數據
// - Channel 可以用於 goroutine 之間的通信和同步

// sum 函數計算切片 s 中所有數字的和，然後將結果發送到 channel c
func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	// 將計算結果發送到 channel
	// 如果 channel 的接收端還沒準備好，這裡會阻塞等待
	c <- sum // send sum to c
}

func RunConcurrency02() {
	// 定義一個包含 6 個數字的切片
	s := []int{7, 2, 8, -9, 4, 0}

	// 創建一個 int 類型的 channel
	// make(chan int) 創建一個無緩衝的 channel（unbuffered channel）
	// 無緩衝 channel 要求發送和接收必須同時準備好，否則會阻塞
	c := make(chan int)

	// 啟動第一個 goroutine，計算切片的前半部分 [7, 2, 8]
	// s[:len(s)/2] 是 s[0:3]，即 [7, 2, 8]
	go sum(s[:len(s)/2], c)

	// 啟動第二個 goroutine，計算切片的後半部分 [-9, 4, 0]
	// s[len(s)/2:] 是 s[3:6]，即 [-9, 4, 0]
	go sum(s[len(s)/2:], c)

	// 從 channel 接收兩個值
	// 這兩個接收操作會阻塞，直到兩個 goroutine 都完成計算並發送結果
	// x 和 y 的順序取決於哪個 goroutine 先完成（可能是任意順序）
	x, y := <-c, <-c // receive from c

	// 打印兩個部分的和以及總和
	// 例如：x=17 (7+2+8), y=-5 (-9+4+0), 總和=12
	fmt.Println(x, y, x+y)

	// 執行流程說明：
	// 1. 主 goroutine 創建 channel 和切片
	// 2. 啟動兩個 goroutine 並發計算不同部分的總和
	// 3. 主 goroutine 在 <-c 處阻塞，等待第一個結果
	// 4. 第一個 goroutine 完成計算，發送結果到 channel，主 goroutine 接收 x
	// 5. 主 goroutine 在第二個 <-c 處阻塞，等待第二個結果
	// 6. 第二個 goroutine 完成計算，發送結果到 channel，主 goroutine 接收 y
	// 7. 計算並打印最終結果
	//
	// 這個例子展示了如何使用 channel 來：
	// - 在 goroutine 之間傳遞數據
	// - 同步多個 goroutine 的執行
	// - 等待所有 goroutine 完成後再繼續
}

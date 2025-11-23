// https://go.dev/tour/concurrency/4
package concurrency

import (
	"fmt"
)

// Channel 的關閉（Close）和 Range 循環
//
// 1. 關閉 Channel：
//    - 發送者可以關閉 channel 來表示不會再發送更多值
//    - 使用 close(ch) 來關閉 channel
//
// 2. 檢測 Channel 是否關閉：
//    接收者可以通過接收表達式的第二個參數來檢測 channel 是否已關閉：
//      v, ok := <-ch
//    如果 channel 已關閉且沒有更多值可接收，ok 為 false
//
// 3. Range 循環：
//    for i := range c 會重複從 channel 接收值，直到 channel 被關閉
//
// 重要注意事項：
// ! 只有發送者應該關閉 channel，接收者永遠不應該關閉
// ! 在已關閉的 channel 上發送會導致 panic
// ! Channel 不像文件，通常不需要關閉它們
//    只有在接收者必須知道不會再有更多值時才需要關閉（例如終止 range 循環）

// fibonacci 函數生成斐波那契數列的前 n 個數字，並發送到 channel
func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for range n {
		c <- x // 發送斐波那契數到 channel
		x, y = y, x+y
	}
	// 發送完所有值後，關閉 channel
	// 這告訴接收者不會再有更多值了
	close(c)
}

func RunConcurrency04() {
	fmt.Println("=== 示例 1: 使用 Range 循環接收值 ===")

	// 創建一個緩衝區大小為 10 的 channel
	c := make(chan int, 10)

	// 啟動 goroutine 生成斐波那契數列
	// cap(c) 返回 channel 的容量（10）
	go fibonacci(cap(c), c)

	// 使用 range 循環從 channel 接收值
	// range 會自動接收值，直到 channel 被關閉
	// 當 channel 關閉且所有值都被接收後，循環會自動結束
	for i := range c {
		fmt.Println(i)
	}

	fmt.Println()

	fmt.Println("=== 示例 2: 手動檢測 Channel 是否關閉 ===")

	ch := make(chan int, 3)
	ch <- 1
	ch <- 2
	ch <- 3
	close(ch) // 關閉 channel

	// 手動接收並檢測 channel 是否關閉
	for {
		v, ok := <-ch
		if !ok {
			// ok 為 false 表示 channel 已關閉且沒有更多值
			fmt.Println("Channel 已關閉，沒有更多值")
			break
		}
		fmt.Println("接收值:", v)
	}

	fmt.Println()

	fmt.Println("=== 示例 3: 從已關閉的 Channel 接收 ===")

	ch2 := make(chan int, 2)
	ch2 <- 10
	ch2 <- 20
	close(ch2)

	// 從已關閉的 channel 接收值是可以的
	// 會返回零值和 false
	v1, ok1 := <-ch2
	fmt.Printf("值: %d, 是否還有值: %v\n", v1, ok1) // 值: 10, 是否還有值: true

	v2, ok2 := <-ch2
	fmt.Printf("值: %d, 是否還有值: %v\n", v2, ok2) // 值: 20, 是否還有值: true

	v3, ok3 := <-ch2
	fmt.Printf("值: %d, 是否還有值: %v\n", v3, ok3) // 值: 0, 是否還有值: false（channel 已關閉）

	fmt.Println()

	fmt.Println("=== 示例 4: 在已關閉的 Channel 上發送會導致 Panic ===")

	ch3 := make(chan int, 2)
	ch3 <- 1
	close(ch3)

	// 嘗試在已關閉的 channel 上發送會導致 panic
	// 取消註釋下面這行會導致程序崩潰：
	// ch3 <- 2 // panic: send on closed channel

	fmt.Println("注意：在已關閉的 channel 上發送會導致 panic")

	fmt.Println()

	fmt.Println("=== 總結 ===")
	fmt.Println("1. 只有發送者應該關閉 channel")
	fmt.Println("2. 使用 close(ch) 關閉 channel")
	fmt.Println("3. 使用 v, ok := <-ch 檢測 channel 是否關閉")
	fmt.Println("4. 使用 for i := range c 循環接收值直到 channel 關閉")
	fmt.Println("5. 在已關閉的 channel 上發送會導致 panic")
	fmt.Println("6. 從已關閉的 channel 接收會返回零值和 false")
}

// https://go.dev/tour/concurrency/6
package concurrency

import (
	"fmt"
	"time"
)

// Select 語句中的 Default Case
//
// default case 在 select 語句中的作用：
// - 如果沒有其他 case 準備好（所有 case 都阻塞），則執行 default case
// - 使用 default case 可以實現非阻塞的發送或接收操作
//
// 語法示例：
//   select {
//   case i := <-c:
//       // 使用 i
//   default:
//       // 從 c 接收會阻塞，所以執行 default
//   }
//
// 重要特性：
// - default case 使得 select 不會阻塞
// - 如果沒有 default，select 會一直阻塞直到至少一個 case 可以執行
// - 有 default 時，如果所有 case 都阻塞，立即執行 default 並繼續
//
// 使用場景：
// - 非阻塞的 channel 操作
// - 在等待多個 channel 的同時執行其他工作
// - 實現超時和定期檢查

func RunConcurrency06() {
	// RunConcurrency06Simple()

	fmt.Println("=== 示例：使用 Default Case 實現非阻塞操作 ===")
	fmt.Println("每 100ms 會收到 tick，500ms 後會收到 boom")
	fmt.Println("在沒有事件時，default case 會執行並打印 '.'")
	fmt.Println()

	start := time.Now()

	// time.Tick 返回一個 channel，每 100ms 發送一次時間
	tick := time.Tick(100 * time.Millisecond)

	// time.After 返回一個 channel，500ms 後發送一次時間
	boom := time.After(500 * time.Millisecond)

	// elapsed 函數返回從開始到現在經過的時間
	elapsed := func() time.Duration {
		return time.Since(start).Round(time.Millisecond)
	}

	for {
		select {
		case <-tick:
			// 每 100ms 執行一次（當 tick channel 有值時）
			fmt.Printf("[%6s] tick.\n", elapsed())
		case <-boom:
			// 500ms 後執行一次，然後退出
			fmt.Printf("[%6s] BOOM!\n", elapsed())
			return
		default:
			// 如果 tick 和 boom 都沒有準備好（都阻塞）
			// 立即執行 default case，不會阻塞
			// 這使得程序可以在等待事件的同時執行其他工作
			fmt.Printf("[%6s]     .\n", elapsed())
			time.Sleep(50 * time.Millisecond)
		}
		// 注意：如果沒有 default case，select 會阻塞直到 tick 或 boom 有值
		// 有了 default case，程序會持續運行，在沒有事件時執行 default
	}

	// 執行流程說明：
	// 1. 程序開始，tick 和 boom 都還沒準備好
	// 2. select 執行 default case，打印 '.'，等待 50ms
	// 3. 重複步驟 2，直到 100ms 後 tick 有值
	// 4. 當 tick 有值時，執行 tick case，打印 "tick."
	// 5. 繼續循環，在 tick 之間執行 default case
	// 6. 500ms 後，boom 有值，執行 boom case，打印 "BOOM!" 並退出
}

// RunConcurrency06Simple 簡單示例：非阻塞的 channel 操作
func RunConcurrency06Simple() {
	fmt.Println("=== 簡單示例：非阻塞的 Channel 操作 ===")

	c := make(chan int)

	// 嘗試非阻塞地接收
	select {
	case i := <-c:
		// 如果 channel 有值，接收並使用
		fmt.Println("接收到值:", i)
	default:
		// 如果 channel 為空（接收會阻塞），執行 default
		fmt.Println("Channel 為空，接收會阻塞，執行 default")
	}

	// 嘗試非阻塞地發送
	select {
	case c <- 42:
		// 如果有接收者在等待或緩衝區有空間，發送成功
		fmt.Println("發送成功")
	default:
		// 如果沒有接收者且緩衝區已滿（發送會阻塞），執行 default
		fmt.Println("Channel 未準備好，發送會阻塞，執行 default")
	}

	fmt.Println()
	fmt.Println("對比：沒有 default 的 select 會阻塞")
	// 如果沒有 default，select 會一直阻塞直到可以執行
	// 取消註釋下面這行會導致程序永遠阻塞：
	// select {
	// case i := <-c:
	// 	fmt.Println("接收到:", i)
	// }
}

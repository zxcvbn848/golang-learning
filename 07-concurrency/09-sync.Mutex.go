// https://go.dev/tour/concurrency/9
package concurrency

import (
	"fmt"
	"sync"
	"time"
)

// 互斥鎖（Mutex）和同步
//
// 我們已經看到 channel 在 goroutine 之間的通信非常有用
// 但如果我們不需要通信呢？如果我們只是想確保同一時間只有一個 goroutine
// 可以訪問一個變量以避免衝突呢？
//
// 這個概念稱為互斥（mutual exclusion），提供這種功能的數據結構
// 的傳統名稱是 mutex（互斥鎖）
//
// Go 的標準庫通過 sync.Mutex 及其兩個方法提供互斥：
//   - Lock()   : 獲取鎖，如果鎖已被其他 goroutine 持有，則阻塞等待
//   - Unlock() : 釋放鎖，允許其他等待的 goroutine 獲取鎖
//
// 我們可以通過在代碼塊周圍調用 Lock 和 Unlock 來定義一個互斥執行的代碼塊
// 也可以使用 defer 來確保 mutex 會被解鎖（如 Value 方法所示）

// SafeCounter 是一個可以在並發環境下安全使用的計數器
// 使用 sync.Mutex 來保護共享的 map
type SafeCounter struct {
	mu sync.Mutex // 互斥鎖，保護下面的 map
	v  map[string]int
}

// Inc 遞增指定鍵的計數器值
// 使用互斥鎖確保同一時間只有一個 goroutine 可以修改 map
func (c *SafeCounter) Inc(key string) {
	c.mu.Lock() // 獲取鎖，如果鎖已被持有，則阻塞等待
	// 鎖定後，只有一個 goroutine 可以訪問 map c.v
	// 這確保了並發安全，避免了數據競爭（data race）
	c.v[key]++
	c.mu.Unlock() // 釋放鎖，允許其他 goroutine 獲取鎖
}

// Value 返回指定鍵的計數器當前值
// 使用 defer 確保即使函數提前返回，鎖也會被釋放
func (c *SafeCounter) Value(key string) int {
	c.mu.Lock()         // 獲取鎖
	defer c.mu.Unlock() // 使用 defer 確保函數返回時自動釋放鎖
	// defer 的好處：
	// - 即使函數中有多個 return 語句，鎖也會被正確釋放
	// - 即使發生 panic，鎖也會被釋放（避免死鎖）
	// - 代碼更簡潔，不需要記住在所有返回點都調用 Unlock
	return c.v[key]
}

func RunConcurrency09() {
	fmt.Println("=== 示例：使用 Mutex 保護共享數據 ===")
	fmt.Println("啟動 1000 個 goroutine 同時遞增計數器")
	fmt.Println()

	// 創建一個 SafeCounter 實例
	c := SafeCounter{v: make(map[string]int)}

	// 啟動 1000 個 goroutine，每個都嘗試遞增同一個鍵的值
	// 如果沒有互斥鎖保護，這會導致數據競爭和不確定的結果
	for range 1000 {
		go c.Inc("somekey")
	}

	// 等待所有 goroutine 完成
	// 注意：在實際應用中，應該使用 sync.WaitGroup 而不是 time.Sleep
	time.Sleep(time.Second)

	// 讀取最終值
	// 由於使用了互斥鎖，結果應該是 1000（每個 goroutine 遞增一次）
	fmt.Printf("最終計數值: %d (期望: 1000)\n", c.Value("somekey"))

	fmt.Println()
	fmt.Println("=== 對比：沒有 Mutex 的情況 ===")
	fmt.Println("如果沒有互斥鎖保護，多個 goroutine 同時修改 map 會導致：")
	fmt.Println("1. 數據競爭（data race）")
	fmt.Println("2. 不確定的結果（可能少於 1000）")
	fmt.Println("3. 程序可能崩潰或產生錯誤的數據")

	fmt.Println()
	fmt.Println("=== Mutex 使用要點 ===")
	fmt.Println("1. 在訪問共享資源前調用 Lock()")
	fmt.Println("2. 訪問完成後調用 Unlock()")
	fmt.Println("3. 使用 defer Unlock() 可以確保鎖一定會被釋放")
	fmt.Println("4. 鎖的持有時間應該盡可能短，避免影響並發性能")
	fmt.Println("5. 不要忘記解鎖，否則會導致死鎖")
}

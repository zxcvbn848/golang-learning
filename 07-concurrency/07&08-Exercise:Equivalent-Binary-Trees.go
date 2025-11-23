// https://go.dev/tour/concurrency/7
// https://go.dev/tour/concurrency/8

package concurrency

import (
	"fmt"

	"golang.org/x/tour/tree"
)

// 練習：等價二叉樹
//
// 問題描述：
// 可能有許多不同的二叉樹存儲相同的值序列
// 例如，兩個不同的二叉樹都可以存儲序列 1, 1, 2, 3, 5, 8, 13
//
// 在大多數語言中，檢查兩個二叉樹是否存儲相同序列的函數相當複雜
// 我們將使用 Go 的並發和 channel 來編寫一個簡單的解決方案
//
// tree.Tree 的結構：
//   type Tree struct {
//       Left  *Tree
//       Value int
//       Right *Tree
//   }

// Walk 函數遍歷樹 t，將樹中的所有值發送到 channel ch
// 使用中序遍歷（in-order traversal）來按順序發送值
// 中序遍歷：左子樹 -> 根節點 -> 右子樹
func Walk(t *tree.Tree, ch chan int) {
	if t == nil {
		return
	}

	// 遞歸遍歷左子樹
	Walk(t.Left, ch)

	// 發送當前節點的值
	ch <- t.Value

	// 遞歸遍歷右子樹
	Walk(t.Right, ch)
}

// walkHelper 是一個輔助函數，用於啟動 Walk 並在完成後關閉 channel
// 這樣接收者可以使用 range 循環來接收所有值
func walkHelper(t *tree.Tree, ch chan int) {
	Walk(t, ch)
	close(ch) // 遍歷完成後關閉 channel
}

// Same 函數判斷兩棵樹 t1 和 t2 是否包含相同的值
// 使用 Walk 函數來獲取兩棵樹的值序列，然後比較它們
func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)

	// 啟動兩個 goroutine 同時遍歷兩棵樹
	go walkHelper(t1, ch1)
	go walkHelper(t2, ch2)

	// 同時從兩個 channel 接收值並比較
	for {
		v1, ok1 := <-ch1
		v2, ok2 := <-ch2

		// 如果兩個 channel 都關閉了，說明所有值都比較完了
		if !ok1 && !ok2 {
			return true // 所有值都相同
		}

		// 如果只有一個 channel 關閉，說明兩棵樹的值數量不同
		if !ok1 || !ok2 {
			return false
		}

		// 如果值不同，兩棵樹不相等
		if v1 != v2 {
			return false
		}
	}
}

func RunConcurrency07() {
	fmt.Println("=== 測試 Walk 函數 ===")

	// tree.New(k) 構造一個隨機結構（但總是排序的）二叉樹
	// 存儲值 k, 2k, 3k, ..., 10k
	// 所以 tree.New(1) 會存儲 1, 2, 3, ..., 10
	t := tree.New(1)

	// 創建一個新的 channel 並啟動 walker
	ch := make(chan int)
	go walkHelper(t, ch)

	// 讀取並打印 10 個值
	// 應該輸出數字 1, 2, 3, ..., 10（按順序）
	fmt.Print("Walk 結果: ")
	for i := range ch {
		fmt.Printf("%d ", i)
	}
	fmt.Println()
	fmt.Println()

	fmt.Println("=== 測試 Same 函數 ===")

	// Same(tree.New(1), tree.New(1)) 應該返回 true
	// 因為兩棵樹都包含相同的值序列（1, 2, 3, ..., 10）
	result1 := Same(tree.New(1), tree.New(1))
	fmt.Printf("Same(tree.New(1), tree.New(1)) = %v (期望: true)\n", result1)

	// Same(tree.New(1), tree.New(2)) 應該返回 false
	// 因為第一棵樹包含 1, 2, 3, ..., 10
	// 而第二棵樹包含 2, 4, 6, ..., 20
	result2 := Same(tree.New(1), tree.New(2))
	fmt.Printf("Same(tree.New(1), tree.New(2)) = %v (期望: false)\n", result2)

	// 額外測試：相同值的不同結構
	fmt.Println()
	fmt.Println("=== 額外測試 ===")

	// 創建兩棵結構不同但值相同的樹
	t1 := tree.New(1)
	t2 := tree.New(1)
	result3 := Same(t1, t2)
	fmt.Printf("Same(兩棵結構不同但值相同的樹) = %v (期望: true)\n", result3)

	fmt.Println()
	fmt.Println("=== 實現說明 ===")
	fmt.Println("1. Walk 函數使用中序遍歷（左-根-右）來按順序發送值")
	fmt.Println("2. Same 函數同時遍歷兩棵樹，逐個比較值")
	fmt.Println("3. 使用 goroutine 和 channel 實現並發遍歷和比較")
	fmt.Println("4. 如果值序列完全相同，返回 true；否則返回 false")
}

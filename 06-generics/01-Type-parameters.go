// https://go.dev/tour/generics/1
package generics

import (
	"fmt"
)

// Go 函數可以使用類型參數（type parameters）來處理多種類型
// 類型參數出現在函數參數之前，用方括號 [] 包圍
//
// Index 函數說明：
// - [T comparable] 是類型參數聲明
//   - T 是類型參數的名稱（可以是任何名稱，通常使用大寫字母）
//   - comparable 是類型約束（constraint），表示 T 必須是可比較的類型
//   - comparable 是 Go 內建的約束，允許使用 == 和 != 運算符
//
// - s []T 表示 s 是類型 T 的切片
// - x T 表示 x 是類型 T 的值
// - 返回值 int 是找到的索引，如果沒找到則返回 -1
//
// 這個函數可以處理任何支持比較操作的類型（如 int, string, float64 等）
func Index[T comparable](s []T, x T) int {
	for i, v := range s {
		// v 和 x 都是類型 T，由於 T 有 comparable 約束
		// 所以我們可以在這裡使用 == 運算符進行比較
		if v == x {
			return i
		}
	}
	return -1
}

func RunGenerics01() {
	// Index 函數可以處理 int 類型的切片
	si := []int{10, 20, 15, -10}
	fmt.Println(Index(si, 15)) // 輸出: 2（15 在索引 2 的位置）

	// Index 函數也可以處理 string 類型的切片
	// 這展示了泛型的強大之處：同一個函數可以處理不同類型的數據
	ss := []string{"foo", "bar", "baz"}
	fmt.Println(Index(ss, "hello")) // 輸出: -1（"hello" 不在切片中）

	// 類型推斷：Go 編譯器可以自動推斷類型參數 T
	// 所以不需要明確指定 Index[int](si, 15)，直接寫 Index(si, 15) 即可
}

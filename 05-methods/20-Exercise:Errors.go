// https://go.dev/tour/methods/20
// Exercise: Errors
// 修改 Sqrt 函數使其返回 error 值，當輸入負數時返回錯誤
package methods

import (
	"fmt"
)

// ErrNegativeSqrt 是一個自訂錯誤類型，用於表示對負數求平方根的錯誤
type ErrNegativeSqrt float64

// Error 實作 error 介面
// 注意：必須先將 e 轉換為 float64，否則 fmt.Sprint(e) 會造成無限循環
// 原因：如果直接使用 fmt.Sprint(e)，fmt 會嘗試呼叫 e.Error()，而 Error() 方法內部又呼叫 fmt.Sprint(e)，形成無限遞迴
func (e ErrNegativeSqrt) Error() string {
	// 轉換為 float64 避免無限循環
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}

// Sqrt 使用牛頓法計算平方根
// 如果輸入為負數，返回 ErrNegativeSqrt 錯誤
func Sqrt(x float64) (float64, error) {
	// 檢查輸入是否為負數
	if x < 0 {
		// 返回錯誤，注意要轉換為 ErrNegativeSqrt 類型
		return 0, ErrNegativeSqrt(x)
	}

	// Newton's method for finding square root
	// Start with an initial guess
	z := 1.0

	// Iterate until we get a good approximation
	for i := 0; i < 10; i++ {
		// Newton's method: z = z - (z*z - x) / (2*z)
		z = z - (z*z-x)/(2*z)
	}

	// 成功時返回結果和 nil error
	return z, nil
}

func RunMethods20() {
	fmt.Println("=== Exercise: Errors ===")
	fmt.Println("測試 Sqrt 函數的錯誤處理")

	// 測試正常情況
	testCases := []float64{2, 4, 7, -2, -4}

	for _, num := range testCases {
		result, err := Sqrt(num)
		if err != nil {
			// 處理錯誤情況
			fmt.Printf("Sqrt(%v) = 錯誤: %v\n", num, err)
		} else {
			// 處理成功情況
			fmt.Printf("Sqrt(%v) = %v\n", num, result)
		}
	}

	fmt.Println("=== 說明 ===")
	fmt.Println("1. ErrNegativeSqrt 實作 error 介面")
	fmt.Println("2. Error() 方法中必須轉換為 float64 避免無限循環")
	fmt.Println("3. Sqrt 函數返回 (float64, error) 兩個值")
	fmt.Println("4. 負數輸入時返回 ErrNegativeSqrt 錯誤")
	fmt.Println("5. 正常情況返回結果和 nil error")
}

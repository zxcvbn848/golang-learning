// https://go.dev/tour/methods/22
// Exercise: Readers
// 實作一個 Reader 類型，它會發送無限的 ASCII 字元 'A' 的流
package methods

import "golang.org/x/tour/reader"

// MyReader 是一個會無限發送 'A' 字元的 Reader
type MyReader struct{}

// Read 實作 io.Reader 介面
// 將提供的 byte slice 填充為 'A' 字元
// 因為是無限流，永遠不會返回 io.EOF 錯誤
func (r MyReader) Read(b []byte) (int, error) {
	// 將整個緩衝區填充為 'A' 字元
	// ASCII 字元 'A' 的值是 65
	for i := range b {
		b[i] = 'A'
	}
	// 返回填充的位元組數（等於緩衝區長度）和 nil 錯誤
	// nil 錯誤表示成功，且因為是無限流，永遠不會返回 io.EOF
	return len(b), nil
}

func RunMethods22() {
	// Validate 函數會測試 MyReader 是否正確實作了 io.Reader 介面
	// 它會讀取一些資料並驗證是否都是 'A' 字元
	reader.Validate(MyReader{}) // OK!
}

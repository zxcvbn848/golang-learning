// https://go.dev/tour/methods/21
// Readers
// io.Reader 介面表示資料流的讀取端，是 Go 中最重要的介面之一
//
// io.Reader 介面定義：
//
//	type Reader interface {
//		Read(p []byte) (n int, err error)
//	}
//
// Read 方法：
//   - 將資料讀取到提供的 byte slice (p) 中
//   - 返回讀取的位元組數 (n) 和錯誤 (err)
//   - 當資料流結束時，返回 io.EOF 錯誤
//
// Go 標準庫中許多型別實作了 io.Reader：
//   - 檔案 (os.File)
//   - 網路連線 (net.Conn)
//   - 壓縮器 (gzip.Reader)
//   - 加密器 (cipher.StreamReader)
//   - 字串 (strings.Reader)
//   - 等等...
package methods

import (
	"fmt"
	"io"
	"strings"
)

func RunMethods21() {
	// strings.NewReader 創建一個從字串讀取資料的 Reader
	// 它實作了 io.Reader 介面
	r := strings.NewReader("Hello, Reader!")

	// 創建一個 8 位元組的緩衝區
	// 每次 Read 調用最多讀取 8 個位元組
	b := make([]byte, 8)

	// 持續讀取直到遇到 io.EOF 錯誤（表示資料流結束）
	for {
		// Read 方法：
		//  - n: 實際讀取的位元組數（可能小於 len(b)）
		//  - err: 錯誤值，nil 表示成功，io.EOF 表示資料流結束
		//  - b: 會被填充讀取到的資料（從 b[0] 到 b[n-1]）
		n, err := r.Read(b)

		// 顯示讀取結果
		// n: 讀取的位元組數
		// err: 錯誤狀態
		// b: 整個緩衝區（包含之前讀取的資料和新的資料）
		fmt.Printf("n = %v err = %v b = %v\n", n, err, b)

		// b[:n] 只包含本次實際讀取的資料
		// 使用 %q 格式化可以清楚看到字串內容
		fmt.Printf("b[:n] = %q\n", b[:n])

		// 檢查是否到達資料流結尾
		// io.EOF 是 io 套件定義的特殊錯誤，表示「End Of File」
		if err == io.EOF {
			break
		}
	}

	/*
		執行結果：
		n = 8 err = <nil> b = [72 101 108 108 111 44 32 82]
		b[:n] = "Hello, R"
		n = 6 err = <nil> b = [101 97 100 101 114 33 32 82]
		b[:n] = "eader!"
		n = 0 err = EOF b = [101 97 100 101 114 33 32 82]
		b[:n] = ""

		說明：
		1. 第一次讀取：讀取前 8 個字元 "Hello, R"
		2. 第二次讀取：讀取剩餘 6 個字元 "eader!"
		3. 第三次讀取：沒有更多資料，返回 n=0 和 io.EOF 錯誤
		4. 注意：b 緩衝區在第二次讀取時仍包含部分舊資料，所以要用 b[:n] 來取得正確的內容
	*/
}

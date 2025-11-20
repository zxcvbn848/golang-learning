// https://go.dev/tour/methods/23
// Exercise: rot13Reader
// 實作一個 rot13Reader 類型，它會將輸入的資料進行 rot13 加密
package methods

import (
	"fmt"
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

// Read 實作 io.Reader 介面，從內部的 reader 讀取資料後套用 rot13 轉換。
// rot13 會將字母位移 13 個字母，非字母字元保持不變。
func (rr *rot13Reader) Read(b []byte) (int, error) {
	n, err := rr.r.Read(b)
	for i := 0; i < n; i++ {
		b[i] = rot13(b[i])
	}
	return n, err
}

// rot13 只轉換英文字母，保持大小寫。
func rot13(c byte) byte {
	switch {
	case 'A' <= c && c <= 'Z':
		return 'A' + (c-'A'+13)%26
	case 'a' <= c && c <= 'z':
		return 'a' + (c-'a'+13)%26
	default:
		return c
	}
}

func RunMethods23() {
	fmt.Println("=== Exercise: rot13Reader ===")
	fmt.Println("以下輸出應該是解碼後的字串：")

	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
	fmt.Println()
}

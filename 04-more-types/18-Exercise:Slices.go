// https://go.dev/tour/moretypes/18
package moreTypes

import (
	"fmt"

	"golang.org/x/tour/pic"
)

func Pic(dx, dy int) [][]uint8 {
	// 創建一個長度為 dy 的 slice，每個元素是一個 []uint8
	picture := make([][]uint8, dy)

	// 使用迴圈為每一行分配一個長度為 dx 的 []uint8
	for y := 0; y < dy; y++ {
		row := make([]uint8, dx)
		for x := 0; x < dx; x++ {
			// 使用 (x+y)/2 來生成灰度值，產生從左上到右下的漸變效果
			// 也可以嘗試其他函數：
			// - x*y：會產生有趣的圖案
			// - x^y：會產生更複雜的圖案
			// - (x^2 + y^2)：會產生圓形漸變
			value := (x + y) / 2
			row[x] = uint8(value)
		}
		picture[y] = row
	}

	return picture
}

func RunMoreTypes18() {
	// 顯示圖片（需要 golang.org/x/tour/pic 套件）
	// 如果套件不可用，可以安裝：go get golang.org/x/tour/pic
	pic.Show(Pic)

	// 也可以打印一些資訊來驗證函數是否正常工作
	result := Pic(5, 5)
	fmt.Println("Pic(5, 5) 的前幾行:")
	for i, row := range result {
		if i < 3 {
			fmt.Printf("  行 %d: %v\n", i, row)
		}
	}
}

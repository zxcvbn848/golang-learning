// https://go.dev/tour/methods/25
// Exercise: Images
// 實作一個自訂的 Image，回傳 color.RGBA 色彩，並用 pic.ShowImage 顯示結果
package methods

import (
	"image"
	"image/color"

	"golang.org/x/tour/pic"
)

// Image 代表一張寬高固定的影像，實作 image.Image 介面
type Image struct {
	width, height int
}

// ColorModel 回傳影像的色彩模型，題目指定 color.RGBAModel
func (img Image) ColorModel() color.Model {
	return color.RGBAModel
}

// Bounds 回傳影像範圍，使用 image.Rect(0, 0, w, h)
func (img Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, img.width, img.height)
}

// At 計算指定像素的色彩
// 這裡沿用先前圖片產生器的概念：用 x、y 計算出數值 v
// 題目指定 color.RGBA{v, v, 255, 255}
func (img Image) At(x, y int) color.Color {
	v := uint8((x ^ y) % 256) // x XOR y，範圍 0-255
	return color.RGBA{v, v, 255, 255}
}

func RunMethods25() {
	// 建立 256x256 的影像並顯示
	m := Image{width: 256, height: 256}
	pic.ShowImage(m)
}

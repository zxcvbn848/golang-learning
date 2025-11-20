// https://go.dev/tour/methods/24
// Images
// Go 的 image 套件定義了 Image 介面：
//
//	type Image interface {
//		ColorModel() color.Model
//		Bounds() Rectangle
//		At(x, y int) color.Color
//	}
//
// - ColorModel 回傳影像使用的色彩模型（例如 color.RGBAModel）
// - Bounds 回傳影像的邊界範圍（左上角、右下角）
// - At 回傳指定座標的色彩，型別為 color.Color（常見實作 color.RGBA）
//
// image.NewRGBA 會建立一個滿足 Image 介面的結構，
// 內部使用 color.RGBA 類型與 color.RGBAModel 來儲存像素。
package methods

import (
	"fmt"
	"image"
)

func RunMethods24() {
	// 建立一個 100x100 的 RGBA 影像，image.Rect 會建立對應的 Rectangle。
	m := image.NewRGBA(image.Rect(0, 0, 100, 100))

	// Bounds() 符合 Image 介面需求，回傳左上(0,0)到右下(100,100) 的矩形。
	fmt.Println("Bounds:", m.Bounds())

	// At(x, y) 取得指定像素的色彩。RGBA() 展開為 16-bit 的 RGBA 成分與 Alpha。
	r, g, b, a := m.At(0, 0).RGBA()
	fmt.Printf("Pixel(0,0) RGBA: R=%d G=%d B=%d A=%d\n", r, g, b, a)
}

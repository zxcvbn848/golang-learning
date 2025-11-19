// https://go.dev/tour/methods/17
// Stringers
// Stringer 是 fmt 套件定義的最常見的 interface 之一
//
//	type Stringer interface {
//	    String() string
//	}
//
// Stringer 是一個可以將自己描述為字串的型別
// fmt 套件（以及許多其他套件）會尋找這個 interface 來列印值
// 當一個型別實作了 String() 方法，fmt.Println 等函數會自動呼叫它
package methods

import (
	"fmt"
)

type Person struct {
	Name string
	Age  int
}

// Person 實作了 Stringer interface
// 當使用 fmt.Println 等函數列印 Person 時，會自動呼叫這個方法
func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

func RunMethods17() {
	a := Person{"Arthur Dent", 42}
	z := Person{"Zaphod Beeblebrox", 9001}
	// fmt.Println 會自動呼叫 Person 的 String() 方法
	// 輸出：Arthur Dent (42 years) Zaphod Beeblebrox (9001 years)
	fmt.Println(a, z)
}

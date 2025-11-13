// https://go.dev/tour/basics/5
package basics

import "fmt"

func add(x, y, z int) int {
	return x + y + z
}

func RunBasics05() {
	fmt.Println(add(42, 13, 50))
}

// https://go.dev/tour/basics/4
package basics

import "fmt"

func add(x int, y int, z int) int {
	return x + y + z
}

func RunBasics04() {
	fmt.Println(add(42, 13, 15))
}

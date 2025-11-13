// https://go.dev/tour/basics/6
package basics

import "fmt"

func swap(x, y, z string) (string, string, string) {
	return z, x, y
}

func RunBasics06() {
	a, b, c := swap("test", "hello", "world")
	fmt.Println(a, b, c)
}

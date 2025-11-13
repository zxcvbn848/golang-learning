// https://go.dev/tour/basics/7
package basics

import "fmt"

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func RunBasics07() {
	fmt.Println(split(17))
}

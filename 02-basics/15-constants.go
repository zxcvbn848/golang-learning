// https://go.dev/tour/basics/15
package basics

import (
	"fmt"
)

const Pi = 3.14

func RunBasics15() {
	const World = "世界"
	fmt.Println("Hello", World)
	fmt.Println("Happy", Pi, "Day")

	const Truth = true
	fmt.Println("Go rules?", Truth)
}

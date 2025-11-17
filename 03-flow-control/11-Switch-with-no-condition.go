// https://go.dev/tour/flowcontrol/11
package flowControl

import (
	"fmt"
	"time"
)

func RunFlowControl11() {
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
}

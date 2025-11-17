// https://go.dev/tour/flowcontrol/10
package flowControl

import (
	"fmt"
	"time"
)

func RunFlowControl10() {
	// 彩蛋：2009-11-10 23:00:00 UTC 是 Go 誕生之日
	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}
}

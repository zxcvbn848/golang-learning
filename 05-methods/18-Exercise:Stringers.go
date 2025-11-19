// https://go.dev/tour/methods/18
package methods

import (
	"fmt"
)

type IPAddr [4]byte

// String 實作 fmt.Stringer 介面，將 4 個位元組格式化成 dotted quad。
// fmt 系列函式（Println、Printf 等）會自動呼叫此方法來取得字串表示。
func (ip IPAddr) String() string {
	return fmt.Sprintf("%d.%d.%d.%d", ip[0], ip[1], ip[2], ip[3])
}

func RunMethods18() {
	// hosts 模擬主機名稱與對應 IP 的查詢表。
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	// 使用 fmt.Printf 時會自動呼叫 IPAddr.String()，因此輸出為「a.b.c.d」格式。
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}

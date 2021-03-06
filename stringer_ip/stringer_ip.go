/***
@desc  go 切割 ip
@问题  这里的 %v 是什么意思~~

*/

package stringer_ip

import "fmt"

type IPAddr [4]byte

func (ip IPAddr) String() string {
	return fmt.Sprintf("%d.%d.%d.%d", ip[0], ip[1], ip[2], ip[3])
}

func main() {
	addrs := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1}, // [127 0 0 1]
		"microsoft": {9, 9, 9, 9},   //  [9,9,9,9]
	}
	//fmt.Println(addrs)
	for n, a := range addrs {
		fmt.Printf("%v: %v\n", n, a)
	}
}

package main

import (
	"fmt"
	"strconv"
)

type IPAddr [4]byte

// TODO: Add a "String() string" method to IPAddr.
func (ipaddr *IPAddr) String() string {
	str := ""
	for i, v := range ipaddr {
		Vint := strconv.Itoa(int(v))
		if i == len(ipaddr)-1 { // 如果是最后一个数字，不用加上 .
			str = str + string(Vint)
			// break   这句写不写都行
		} else {
			str = str + string(Vint) + "."
		}

	}
	return str

}
func main() {
	hosts := map[string]*IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
	a := &IPAddr{1, 2, 3, 4}
	fmt.Println(a)
}

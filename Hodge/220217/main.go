package main

import (
	"fmt"
	"net"
)

func main() {
	a := true
	fmt.Printf("%v\n", a)
}

func GetHostIP() string {
	var results []string
	ifaces, err := net.Interfaces()
	if err != nil {
		fmt.Println(err)
	}
	var addrs []net.Addr
	for _, i := range ifaces {
		if i.HardwareAddr.String() == "" {
			continue
		}
		addrs, err = i.Addrs()
		if err != nil {
			fmt.Println(err)
		}
		for _, addr := range addrs {
			if ip := addr.(*net.IPNet).IP.To4(); ip != nil {
				if ip[0] != 127 && ip[0] != 169 {
					results = append(results, ip.String())
				}
			}
		}
	}
	return results[len(results)-1]
}

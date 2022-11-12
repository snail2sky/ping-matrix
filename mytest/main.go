package main

import (
	"fmt"
	"net"
	"os"
)

func GetLocalIP() string {
	var ip string
	addresses, err := net.InterfaceAddrs()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	for _, address := range addresses {
		// 检查ip地址判断是否回环地址
		if ipNet, ok := address.(*net.IPNet); ok && !ipNet.IP.IsLoopback() {
			if ipNet.IP.To4() != nil {
				ip = ipNet.IP.String()
				fmt.Println(ip)
			}
		}
	}
	return ip
}

func main() {
	fmt.Println(GetLocalIP())
}

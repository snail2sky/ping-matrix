package tools

import (
	"log"
	"net"
	"strings"
)

func GetLocalIPByTCPConnect(addr string) string {
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		log.Fatalln(err)
	}
	err = conn.Close()
	if err != nil {
		return ""
	}
	ip := strings.Split(conn.LocalAddr().String(), ":")[0]
	return ip
}

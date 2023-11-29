package main

import (
	"net"
	"strconv"
	"time"
)

func ScanPort(protocal, hostname string, port int) bool {
	address := hostname + ":" + strconv.Itoa(port)
	conn, err := net.DialTimeout(protocal, address, 60*time.Second)
	if err != nil {
		return false
	}
	defer conn.Close()
	return true
}

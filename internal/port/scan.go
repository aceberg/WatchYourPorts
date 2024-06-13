package port

import (
	"fmt"
	"net"
	"strconv"
	"time"
)

func isOpen(host string, port int) bool {

	timeout := 3 * time.Second
	target := fmt.Sprintf("%s:%d", host, port)

	conn, err := net.DialTimeout("tcp", target, timeout)
	if err != nil {
		return false
	}

	if conn != nil {
		conn.Close()
		return true
	}

	return false
}

// Scan - scan all TCP ports of a host
func Scan(host, beginStr, endStr string) (ports []int) {

	begin, _ := strconv.Atoi(beginStr)
	end, _ := strconv.Atoi(endStr)

	for i := begin; i < end+1; i++ {
		if isOpen(host, i) {
			ports = append(ports, i)
		}
	}

	return ports
}

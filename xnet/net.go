package xnet

import (
	"fmt"
	"net"
	"strings"
)

// HostPort format addr and port suitable for dial
func HostPort(addr string, port interface{}) string {
	host := addr
	if strings.Count(addr, ":") > 0 {
		host = fmt.Sprintf("[%s]", addr)
	}
	// when port is blank or 0, host is a queue name
	if v, ok := port.(string); ok && v == "" {
		return host
	} else if v, ok := port.(int); ok && v == 0 && net.ParseIP(host) == nil {
		return host
	}

	return fmt.Sprintf("%s:%v", host, port)
}

// // GetLocalMainIP ...
// func GetLocalMainIP() (string, error) {
// 	// UDP Connect, no handshake
// 	conn, err := net.Dial("udp", "8.8.8.8:53")
// 	if err != nil {
// 		return "", err
// 	}
// 	defer conn.Close()
// 	localAddr := conn.LocalAddr().(*net.UDPAddr)

// 	return localAddr.IP.String(), nil
// }

package utils

import (
	"fmt"
	"net"
	"regexp"
	"strconv"
	"time"
)

// ポートが立ち上がっているか確かめる
func IsFoundHost(host string, port uint16) bool {
	target := fmt.Sprintf("%s:%d", host, port)

	_, err := net.DialTimeout("tcp", target, 1*time.Second)
	if err != nil {
		fmt.Printf("%s %v\n", target, err)
		return false
	}

	return true
}

var PATTERN = regexp.MustCompile(`((25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?\.){3})(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)`)

// 自分の近くの可動しているサーバーのアドレスを見つける
func FindNeighbors(myHost string, myPort uint16, startIp uint8, endIp uint8, startPort uint16, endPort uint16) []string {
	address := fmt.Sprintf("%s:%d", myHost, myPort)

	// example myHost(127.0.0.1) -> m [127.0.0.1 127.0.0. 0. 1]
	m := PATTERN.FindStringSubmatch(myHost)
	if m == nil {
		return nil
	}
	prefixHost := m[1]                     // example 127.0.0.
	lastIp, _ := strconv.Atoi(m[len(m)-1]) // example 1
	neighbors := make([]string, 0)

	for port := startPort; port <= endPort; port += 1 {
		for ip := startIp; ip <= endIp; ip += 1 {
			guessHost := fmt.Sprintf("%s%d", prefixHost, lastIp+int(ip))
			guessTarget := fmt.Sprintf("%s:%d", guessHost, port) // example 127.0.0.1:5001

			// アドレスが自分のアドレスと同じではない、かつ稼働しているアドレスを返り値の配列に追加
			if guessTarget != address && IsFoundHost(guessHost, port) {
				neighbors = append(neighbors, guessTarget)
			}
		}
	}

	return neighbors
}

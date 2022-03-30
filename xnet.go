package xnet

import "net"

func CheckIP(s string) bool {
	if _, err := net.ResolveIPAddr("", s); err != nil {
		return false
	}
	return true
}

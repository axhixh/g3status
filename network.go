package main

import (
	"net"
	"strings"
)

func GetIpAddr() *Block {
	addrs, _ := net.InterfaceAddrs()
	for _, addr := range addrs {
		if strings.HasPrefix(addr.String(), "127") {
			continue
		}
		return &Block{"net", addr.Network(), addr.String(), ""}
	}
	return &Block{"net", "default", "No network", ""}
}

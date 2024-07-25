package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: cidr <CIDR-block>")
		return
	}

	cidr := os.Args[1]

	ip, ipnet, err := net.ParseCIDR(cidr)
	if err != nil {
		fmt.Println("Invalid CIDR block:", err)
	}

	fmt.Println("IP Address:", ip)
	fmt.Println("Network:", ipnet)
	fmt.Println("Network Address:", ipnet.IP)

	broadcast := calculateBroadcastAddress(ipnet)
	fmt.Println("Broadcast Address:", broadcast)

}

func calculateBroadcastAddress(ipnet *net.IPNet) net.IP {
	ip := ipnet.IP.To4()
	// mask := ipnet.Mask
	broadcast := make(net.IP, len(ip))
	copy(broadcast, ip)
	for i := range broadcast {
		broadcast[i] |= ^ipnet.Mask[i]
	}
	return broadcast
}

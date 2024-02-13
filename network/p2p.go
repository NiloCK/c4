package network

import (
	"fmt"
	"net"
	"time"
)

const broadcastAddr = "192.168.2.255:4444"
const listenAddr = ":4444"

func LookupBroadcastAddr() {
	ifaces, err := net.Interfaces()
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, iface := range ifaces {
		if iface.Flags&net.FlagUp == 0 {
			continue // interface down
		}
		if iface.Flags&net.FlagBroadcast == 0 {
			continue // interface does not support broadcast
		}

		addrs, err := iface.Addrs()
		if err != nil {
			fmt.Println(err)
			continue
		}

		for _, addr := range addrs {
			ipNet, ok := addr.(*net.IPNet)
			if !ok {
				continue
			}

			ip := ipNet.IP.To4()
			if ip == nil {
				continue // not an ipv4 address
			}

			mask := ipNet.Mask
			broadcast := net.IP(make([]byte, 4))
			for i := range ip {
				broadcast[i] = ip[i] | ^mask[i]
			}
			fmt.Println("Interface:", iface.Name, "Broadcast Address:", broadcast)
		}
	}
}

func Broadcast() {
	// ping the broadcast address
	conn, err := net.Dial("udp", broadcastAddr)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	for {

		_, err := conn.Write([]byte("c4"))

		if err != nil {
			fmt.Println("broadcast write error: ", err)
			return
		}
		fmt.Println("broadcast sent")
		time.Sleep(1 * time.Second)
	}

}

func Listen() {
	addr, err := net.ResolveUDPAddr("udp", listenAddr)
	if err != nil {
		fmt.Println("ResolveUDPAddr error:", err)
		return
	}

	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		fmt.Println("ListenUDP error:", err)
		return
	}
	defer conn.Close()

	for {
		buffer := make([]byte, 1024)
		n, src, err := conn.ReadFromUDP(buffer)
		if err != nil {
			fmt.Println("ReadFromUDP error:", err)
			continue
		}
		fmt.Printf("Received message: %s from %s\n", string(buffer[:n]), src)
	}
}

package main

import (
	"fmt"
	"net"
	"os"
	"strconv"
	"time"
)

func main() {
	if len(os.Args) < 4 {
		fmt.Println("Usage: go run main.go <host> <startPort> <endPort>")
		return
	}

	host := os.Args[1]
	startPort, _ := strconv.Atoi(os.Args[2])
	endPort, _ := strconv.Atoi(os.Args[3])

	fmt.Printf("🔍 Scanning host %s from port %d to %d...\n", host, startPort, endPort)

	for port := startPort; port <= endPort; port++ {
		address := fmt.Sprintf("%s:%d", host, port)
		conn, err := net.DialTimeout("tcp", address, 500*time.Millisecond)

		if err == nil {
			fmt.Printf("✅ Port %d OPEN\n", port)
			conn.Close()
		}
	}
	fmt.Println("✅ Scan complete.")
}

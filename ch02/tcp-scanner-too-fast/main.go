package main

import (
	"fmt"
	"net"
)

func main() {
	for i := 1; i < 1024; i++ {
		go func(j int) {
			addres := fmt.Sprintf("scanme.nmap.org:%d", i)
			conn, err := net.Dial("tcp", addres)
			if err != nil {
				fmt.Printf("Connection FAILED to the port or PORT CLOSED: %d\n", i)
				return
			}
			conn.Close()
			fmt.Printf("Connection SUCCESSFUL to the port: %d\n", i)
		}(i)
	}
}

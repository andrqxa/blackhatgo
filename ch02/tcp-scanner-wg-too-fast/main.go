package main

import (
	"fmt"
	"net"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	for i := 1; i < 1024; i++ {
		wg.Add(1)
		go func(j int) {
			defer wg.Done()
			addres := fmt.Sprintf("scanme.nmap.org:%d", j)
			conn, err := net.Dial("tcp", addres)
			if err != nil {
				fmt.Printf("Connection FAILED to the port or PORT CLOSED: %d\n", j)
				return
			}
			conn.Close()
			fmt.Printf("Connection SUCCESSFUL to the port: %d\n", j)

		}(i)
	}
	wg.Wait()
}

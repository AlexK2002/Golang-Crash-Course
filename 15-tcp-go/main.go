package main

import (
	"fmt"
	"net"
	"sync"
)

func main() {
	//wireshark tcp.port == 80 (Browser request: http.host matches "scanme\.nmap\.org")
	// _, err := net.Dial("tcp", "scanme.nmap.org:80")

	// if err == nil {
	// 	fmt.Println(err)
	// }

	// Common 1024 ports
	// for i := 22; i <= 22; i++ {
	// 	address := fmt.Sprintf("scanme.nmap.org:%d", i)

	// 	// Sending SYN request
	// 	conn, err := net.Dial("tcp", address)

	// 	if err != nil {
	// 		//port is closed
	// 		//Response RST/ACK
	// 		//dial tcp 45.33.32.156:20: connectex: No connection could be made because the target machine actively refused it.
	// 		continue
	// 	}
	// 	conn.Close()

	// 	// Response SYN/ACK
	// 	fmt.Printf("%d Open %v\n", i, err)
	// }

	//Better and Faster Way
	var wg sync.WaitGroup
	for i := 1024; i <= 65335; i++ {
		wg.Add(1)
		go func(j int) {
			defer wg.Done()

			address := fmt.Sprintf("127.0.0.1:%d", j)
			conn, err := net.Dial("tcp", address)
			if err != nil {
				return
			}
			conn.Close()
			fmt.Printf("+ %d open\n", j)
		}(i)
	}
	wg.Wait()
}

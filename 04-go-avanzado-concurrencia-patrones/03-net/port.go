package main

import (
	"flag"
	"fmt"
	"net"
	"sync"
)

var site = flag.String("site", "scanme.nmap.org", "to scan ports")

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 65535; i++ {
		flag.Parse()
		wg.Add(1)
		go func(port int) {
			defer wg.Done()
			conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", *site, port))
			if err != nil {
				return
			}
			err = conn.Close()
			fmt.Printf("Port %d esta abierto\n", port)
		}(i)
	}
	wg.Wait()
}

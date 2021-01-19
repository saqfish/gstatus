package main

import (
	"bufio"
	"fmt"
	"net"
)

func getBannerstr(c net.Conn) {
	fmt.Printf("Client connected [%s]\n", c.RemoteAddr().Network())
	defer c.Close()
	cd := bufio.NewReaderSize(c, 256)
	line, _, err := cd.ReadLine()
	if err == nil {
		fmt.Printf("Client sent %s\n", line)
		s := string(line)
		chans[bannerPos] <- s
	}
}

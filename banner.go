package main

import (
	"bufio"
	"net"
)

func getBannerstr(c net.Conn) {
	defer c.Close()
	cd := bufio.NewReaderSize(c, 256)
	line, _, err := cd.ReadLine()
	if err == nil {
		s := string(line)
		chans[bannerPos] <- s
	}
}

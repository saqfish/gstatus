package main

import (
	"fmt"
	"net"
	"os"
	"strings"
	"time"
)

var runs int = 0
var chans [10]chan string
var lst string

func send(s string) {
	fmt.Printf("Sending %s\n", s)
	c, err := net.Dial("unix", gsock)
	defer c.Close()
	if err != nil {
		fmt.Println("socket error")
		os.Exit(1)

	} else {
		_, werr := c.Write([]byte(s))
		if werr != nil {
			fmt.Println("socket write error")
			os.Exit(1)
		}
	}
}

func setroot(c ...string) {
	s := strings.Join(c, ",")
	if c == nil {
		fmt.Println(s)
		send(lst)
	} else {

		if s != lst {
			fmt.Println(s)
			send(s)
			runs++
			stats()
		}
		lst = s
	}
}

func init() {
	for i := range chans {
		chans[i] = make(chan string)
	}
	if err := os.RemoveAll(sock); err != nil {
		os.Exit(1)
	}
	go func() {
		s, err := net.Listen("unix", sock)
		defer s.Close()
		if err != nil {
			os.Exit(1)
		}
		for {
			c, _ := s.Accept()
			go getBannerstr(c)
		}
	}()
}

func run(b ...string) {
	for {
		select {
		case z := <-chans[bannerPos]:
			b[bannerPos] = z
		case z := <-chans[datePos]:
			b[datePos] = z
		case <-time.After(time.Second):
			break
		}
		setroot(b...)

	}
}

func main() {
	bannerCell, dateCell := "00?", "00?"
	go str(2, 1, chans[bannerPos])
	go date(60*1000, 1, chans[datePos])
	run(bannerCell, dateCell)
}

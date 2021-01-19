package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"runtime"
	"strings"
	"time"
)

const (
	red    int = 1
	yellow int = 2
	green  int = 3
)

var runs int = 0
var lst string

var ch3 chan string = make(chan string)
var ch4 chan string = make(chan string)

func send(s string) {
	fmt.Printf("Sending %s\n", s)
	c, err := net.Dial("unix", "/tmp/gsock.sock")
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

func stats() {
	var mem runtime.MemStats
	runtime.ReadMemStats(&mem)
	fmt.Printf("alloc [%v] \t heapAlloc [%v] \n", mem.Alloc, mem.HeapAlloc)
	fmt.Println("Routines: ", runtime.NumGoroutine())
}

const sock = "/tmp/gstatus.sock"

func getBannerstr(c net.Conn) {
	fmt.Printf("Client connected [%s]\n", c.RemoteAddr().Network())
	defer c.Close()
	cd := bufio.NewReaderSize(c, 24)
	line, _, err := cd.ReadLine()
	if err == nil {
		fmt.Printf("Client sent %s\n", line)
		s := string(line)
		ch3 <- s
	}
}

func init() {
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

func banner(s string) {
	go str(2, 1, ch3)
	ch3 <- s
}

func main() {
	m, d := "10?", "10?"
	banner("first")
	go date(60*1000, 1, ch4)

	for {
		select {
		case z := <-ch3:
			m = z
		case z := <-ch4:
			d = z
		case <-time.After(time.Second):
			break
		}
		setroot(d, m)

	}
}

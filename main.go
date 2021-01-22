package main

import (
	"fmt"
	"net"
	"os"
	"strings"
	"time"
)

var runs, tries int
var chans [5]chan string
var lst string

func send(s string) bool {
	c, err := net.Dial("unix", gsock)
	if err != nil {
		fmt.Println("socket error")
		return false
	} else {
		_, werr := c.Write([]byte(s))
		if werr != nil {
			fmt.Println("socket write error")
			return false
		}
	}
	c.Close()
	return true
}

func setroot(c ...string) {
	s := strings.Join(c, ",")
	var sent bool
	if c == nil {
		sent = send(lst)
	} else {
		if s != lst || tries < len(chans) {
			sent = send(s)
			runs++
			stats()
		} else {
			sent = true
		}
		lst = s
	}

	// make 10 attempts to send before quiting
	if !sent {
		if tries == len(chans) {
			os.Exit(1)
		}
		tries++
		fmt.Println("tries: ", tries)
		time.Sleep(1000)
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
		case z := <-chans[cpuPos]:
			b[cpuPos] = z
		case z := <-chans[hourPos]:
			b[hourPos] = z
		case z := <-chans[batteryPos]:
			b[batteryPos] = z
		case z := <-chans[datePos]:
			b[datePos] = z
		case <-time.After(time.Second):
			break
		}
		setroot(b...)

	}
}

func main() {
	bannerCell, hourCell, cpuCell, dateCell, batteryCell := "00?", "00?", "00?", "00?", "00?"
	go date(60, 1, chans[datePos])
	go battery(10, "/sys/class/power_supply/BAT0/capacity", 1, chans[batteryPos])
	go hour(30, 1, chans[hourPos])
	go cpu_perc(4, 1, chans[cpuPos])
	go str(2, 1, chans[bannerPos])
	run(bannerCell, cpuCell, hourCell, dateCell, batteryCell)
}

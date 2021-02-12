package main

import (
	"fmt"
	"net"
	"os"
)

var runs int
var chans [6]chan string

func line(c ...string) {
	var s, b string
	reset := " %{-u}%{F-}%{B-} "
	for n, l := range c {
		if n == 0 {
			b = l
		} else {
			s = fmt.Sprintf("%s%s%s", s, l, reset)
		}
	}
	ls := fmt.Sprintf("%s%%{r}%s", b, s)
	fmt.Println(ls)
	runs++
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
		case z := <-chans[hourPos]:
			b[hourPos] = z
		case z := <-chans[cpuPos]:
			b[cpuPos] = z
		case z := <-chans[ramPos]:
			b[ramPos] = z
		case z := <-chans[batteryPos]:
			b[batteryPos] = z
		case z := <-chans[datePos]:
			b[datePos] = z
		}
		line(b...)
	}
}

func mkline(fg string, bg string, v string) string {
	s := fmt.Sprintf("%%{F%s}%%{B%s}%s", fg, bg, v)
	return s
}
func mkuline(bs string, c string, v string) string {
	s := fmt.Sprintf("%%{U%s}%%{+u}%s", c, v)
	return s
}

func main() {
	var bannerCell, hourCell, cpuCell, ramCell, dateCell, batteryCell string

	bg := "#0a1016"
	go date(60, bg, chans[datePos])
	go battery(10, "/sys/class/power_supply/BAT0/capacity", bg, chans[batteryPos])
	go ramperc(4, bg, chans[ramPos])
	go cpuperc(4, bg, chans[cpuPos])
	go hour(30, bg, chans[hourPos])
	go str(green, bg, chans[bannerPos])

	run(bannerCell, hourCell, cpuCell, ramCell, dateCell, batteryCell)
}

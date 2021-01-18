package main

import (
	"fmt"
	"os/exec"
	"time"
)

const (
	red    int = 1
	yellow int = 2
	green  int = 3
)

var runs int = 0

func main() {
	var one int
	c, m, d := "1?", "1?", "1?"
	for {
		go num(100, 1, &one, &c)
		go str(5000, 2, "Hello there!", &m)
		go date(60*1000, &d)

		s := fmt.Sprintf("%s,%s,%s", d, m, c)
		exec.Command("xsetroot", "-name", s).Run()
		time.Sleep(1000 * time.Millisecond)
		runs++
	}
}

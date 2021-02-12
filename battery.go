package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"
)

func battery(d float64, p string, bg string, s chan string) {
	for {
		f, ferr := os.Open(p)
		if ferr != nil {
			s <- mkline(red, bg, "!!")
			break
		}
		r := bufio.NewReaderSize(f, 1024)
		line, _, lerr := r.ReadLine()
		if lerr != nil {
			s <- mkline(red, bg, "!!")
			break
		}
		v, perr := strconv.ParseInt(string(line), 10, 64)
		if perr != nil {
			s <- mkline(red, bg, "!!")
			break
		}
		cline := fmt.Sprintf("%s%%", line)
		s <- mkline(clrbttry(int(v)), bg, cline)
		f.Close()
		if runs > 0 {
			time.Sleep(time.Duration(d) * time.Second)
		}
	}
}

func clrbttry(i int) string {
	var clr string
	switch {
	case i <= 20:
		clr = red
	case i <= 60 && i > 20:
		clr = yellow
	case i <= 100 && i > 60:
		clr = green
	}
	return clr
}

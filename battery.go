package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"
)

func battery(d float64, p string, inv int, s chan string) {
	for {
		f, ferr := os.Open(p)
		if ferr != nil {
			s <- fmt.Sprintf("%d%d!!", red, inv)
			break
		}
		r := bufio.NewReaderSize(f, 1024)
		line, _, lerr := r.ReadLine()
		if lerr != nil {
			s <- fmt.Sprintf("%d%d!!", red, inv)
			break
		}
		v, perr := strconv.ParseInt(string(line), 10, 64)
		if perr != nil {
			s <- fmt.Sprintf("%d%d!!", red, inv)
			break
		}
		s <- fmt.Sprintf("%d%d%s%%", clrbttry(int(v)), inv, line)
		f.Close()
		if runs > 0 {
			time.Sleep(time.Duration(d) * time.Second)
		}
	}
}

func clrbttry(i int) int {
	var clr int
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

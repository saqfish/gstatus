package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"time"
)

func ramperc(d float64, bg string, s chan string) {
	for {
		total, free, buffers, cached := memstat()

		used := total - free
		extra := buffers + cached
		cal := (used - extra) * 100

		perc := cal / total

		cperc := fmt.Sprintf("%d%%", perc)
		s <- mklblline("RAM", clrram(perc), bg, cperc)
		if runs > 0 {
			time.Sleep(time.Duration(d) * time.Second)
		}
	}
}

func memstat() (int, int, int, int) {
	var total, free, buffers, cached int
	fp, ferr := os.Open("/proc/meminfo")
	if ferr != nil {
		return 0, 0, 0, 0
	}
	r := bufio.NewReader(fp)
	for {
		rl, rerr := r.ReadString('\n')

		fmt.Sscanf(rl, "MemTotal: %d kB\n", &total)
		fmt.Sscanf(rl, "MemFree: %d kB\n", &free)
		fmt.Sscanf(rl, "Buffers: %d kB\n", &buffers)
		fmt.Sscanf(rl, "Cached: %d kB\n", &cached)

		if rerr == io.EOF {
			break
		}
	}
	return total, free, buffers, cached
}

func clrram(i int) string {
	switch {
	case i > 60:
		return red
	case i > 40:
		return yellow
	default:
		return green
	}
}

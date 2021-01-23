package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func cpuperc(d float64, inv int, s chan string) {
	var po, pt int
	var o, t int
	for {
		_, po, pt = stat()
		if runs > 0 {
			time.Sleep(time.Duration(d) * time.Second)
		}

		_, o, t = stat()

		if pt == 0 {
			s <- fmt.Sprintf("%d%d0%%", green, inv)
			continue
		}

		if pt == t {
			s <- fmt.Sprintf("%d%d0%%", green, inv)
			continue
		}

		perc := 100 * (o - po) / (t - pt)

		s <- fmt.Sprintf("%d%d%d%%", clrcpu(perc), inv, perc)
	}
}

const (
	user       int = 0
	nice       int = 1
	system     int = 2
	idle       int = 3
	iowait     int = 4
	irq        int = 5
	softirq    int = 6
	steal      int = 7
	guest      int = 8
	guest_nice int = 9
)

func stat() (int, int, int) {
	var cpu string
	fp, ferr := os.Open("/proc/stat")
	if ferr != nil {
		return 0, 0, 0
	}
	r := bufio.NewReader(fp)
	var vars [10]int
	rl, rerr := r.ReadString('\n')
	fmt.Sscanf(rl, "%s %d %d %d %d %d %d %d %d %d %d", &cpu, &vars[user], &vars[nice], &vars[system], &vars[idle], &vars[iowait], &vars[irq], &vars[softirq], &vars[steal], &vars[guest], &vars[guest_nice])
	if rerr != nil {
		os.Exit(1)
	}

	idle := vars[idle] + vars[iowait]
	other := vars[user] + vars[nice] + vars[system] + vars[irq] + vars[softirq] + vars[steal]
	total := idle + other

	return idle, other, total
}

func clrcpu(i int) int {
	var clr int
	switch {
	case i <= 100 && i > 80:
		clr = red
	case i <= 60 && i > 50:
		clr = yellow
	case i <= 40:
		clr = green
	}
	return clr
}

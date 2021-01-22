package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func pstat(i int, o int, t int) {
	fmt.Println("idle: ", i)
	fmt.Println("other: ", o)
	fmt.Println("total: ", t)
	fmt.Println("---------")
}

func cpu_perc(d float64, inv int, s chan string) {
	var pi, po, pt int
	var i, o, t int
	for {
		pi, po, pt = stat()
		pstat(pi, po, pt)
		if runs > 0 {
			time.Sleep(time.Duration(d) * time.Second)
		}

		i, o, t = stat()
		pstat(i, o, t)

		if pt == 0 {
			s <- fmt.Sprintf("%d%d0%%", green, inv)
			continue
		}

		if pt == t {
			s <- fmt.Sprintf("%d%d0%%", green, inv)
			continue
		}

		perc := 100 * (o - po) / (t - pt)

		s <- fmt.Sprintf("%d%d%d%%", green, inv, perc)
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
		fmt.Println("file error")
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

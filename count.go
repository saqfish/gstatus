package main

import (
	"fmt"
	"time"
)

func count(d float64, c int, bg string, i *int, s chan string) {
	for {
		*i = *i + 1
		cc := fmt.Sprintf("%d", c)
		ci := fmt.Sprintf("%d", *i)
		s <- mkline(cc, bg, ci)
		if runs > 0 {
			time.Sleep(time.Duration(d) * time.Second)
		}
	}
}

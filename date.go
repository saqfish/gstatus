package main

import (
	"time"
)

func date(d float64, bg string, s chan string) {
	for {
		t := time.Now()
		s <- mkline(white, bg, t.Format(time.Kitchen))
		if runs > 0 {
			time.Sleep(time.Duration(d) * time.Second)
		}
	}
}

package main

import (
	"time"
)

func date(d float64, bg string, s chan string) {
	for {
		t := time.Now()
		_, m, _ := t.Clock()
		s <- mkline(clrdte((60 - m)), bg, t.Format(time.Kitchen))
		if runs > 0 {
			time.Sleep(time.Duration(d) * time.Second)
		}
	}
}

func clrdte(i int) string {
	switch {
	case i > 40:
		return green
	case i > 20:
		return yellow
	default:
		return red
	}
}

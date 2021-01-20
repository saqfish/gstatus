package main

import (
	"fmt"
	"time"
)

func hour(d float64, inv int, s chan string) {
	for {
		_, m, _ := time.Now().Clock()
		rem := 60 - m
		clr := clrhr(rem)
		s <- fmt.Sprintf("%d%d%d", clr, inv, rem)
		if runs > 0 {
			time.Sleep(time.Duration(d) * time.Second)
		}
	}
}

func clrhr(i int) int {
	var clr int
	switch {
	case i <= 20:
		clr = red
	case i <= 40 && i > 20:
		clr = yellow
	case i <= 60 && i > 40:
		clr = green
	}
	return clr
}

package main

import (
	"fmt"
	"time"
)

func date(d int, inv int, s chan string) {
	for {
		t := time.Now()
		s <- fmt.Sprintf("%d%d%s", green, inv, t.Format(time.Kitchen))
		if runs > 0 {
			time.Sleep(time.Duration(d) * time.Millisecond)
		}
	}
}

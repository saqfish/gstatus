package main

import (
	"fmt"
	"time"
)

func count(d int, c int, inv int, i *int, s chan string) {
	for {
		*i = *i + 1
		s <- fmt.Sprintf("%d%d%d", c, inv, *i)
		if runs > 0 {
			time.Sleep(time.Duration(d) * time.Millisecond)
		}
	}
}

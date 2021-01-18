package main

import (
	"fmt"
	"time"
)

func date(d int, s *string) {
	if runs > 1 {
		time.Sleep(time.Duration(d) * time.Millisecond)
	}
	t := time.Now()
	*s = fmt.Sprintf("%d%s", green, t.Format(time.Kitchen))
}

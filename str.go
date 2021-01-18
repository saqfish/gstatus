package main

import (
	"fmt"
	"time"
)

func str(d int, c int, w string, s *string) {
	if runs > 1 {
		time.Sleep(time.Duration(d) * time.Millisecond)
	}
	*s = fmt.Sprintf("%d%s", c, w)
}

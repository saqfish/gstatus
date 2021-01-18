package main

import (
	"fmt"
	"time"
)

func num(d int, c int, i *int, s *string) {
	if runs > 1 {
		time.Sleep(time.Duration(d) * time.Millisecond)
	}
	*i = *i + 1
	*s = fmt.Sprintf("%d%d", c, *i)
}

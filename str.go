package main

import "fmt"

func str(c int, inv int, s chan string) {
	var w string
	if runs < 1 {
		w = "Welcome!"
	} else {
		w = <-s
	}
	s <- fmt.Sprintf("%d%d%s", c, inv, w)
}

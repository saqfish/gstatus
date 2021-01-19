package main

import "fmt"

func str(c int, inv int, s chan string) {
	w := <-s
	s <- fmt.Sprintf("%d%d%s", c, inv, w)
}

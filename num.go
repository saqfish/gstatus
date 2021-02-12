package main

import "fmt"

func num(c string, i int, bg string, s chan string) {
	ci := fmt.Sprintf("%d", i)
	s <- mkline(c, bg, ci)
}

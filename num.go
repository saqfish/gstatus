package main

import "fmt"

func num(c int, i int, inv int, s chan string) {
	s <- fmt.Sprintf("%d%d%d", c, inv, i)
}

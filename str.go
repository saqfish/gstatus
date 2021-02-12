package main

func str(c string, bg string, s chan string) {
	if runs > 1 {
		w := <-s
		s <- mkline(c, bg, w)
	}
}

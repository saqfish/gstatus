package main

type Command struct {
	fg  string
	bg  string
	cmd string
	lbl string
}

func shorts(bg string, s chan string) {
	full := []Command{
		{fg: white, bg: bg, cmd: "chromium", lbl: "★"},
		{fg: white, bg: bg, cmd: `st -f "$THEME_TERM_FONT"`, lbl: "☰"},
	}
	var btns string
	for _, btn := range full {
		btns += mkbttn(btn)
	}
	s <- btns
}

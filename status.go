package main

import "fmt"

var runs int

const reset string = "%{-u}%{F-}%{B-}%{O6}"

const (
	datePos    = 4
	batteryPos = 3
	ramPos     = 2
	cpuPos     = 1
	bannerPos  = 0

	white  string = "#FFFFFF"
	red    string = "#ff0000"
	yellow string = "#ffff00"
	green  string = "#00FF00"
)

func line(c ...string) {
	var s, b string
	for n, l := range c {
		if n == 0 {
			b = l
		} else {
			s = fmt.Sprintf("%s%s%s", s, l, reset)
		}
	}
	ls := fmt.Sprintf("%s%%{r}%s", b, s)
	fmt.Println(ls)
	runs++
}

func mkline(fg string, bg string, v string) string {
	s := fmt.Sprintf("%%{F%s}%%{B%s}%s", fg, bg, v)
	return s
}
func mklblline(l string, fg string, bg string, v string) string {
	s := fmt.Sprintf("%s%s%s", l, reset, mkline(fg, bg, v))
	return s
}
func mkuline(bs string, c string, v string) string {
	s := fmt.Sprintf("%%{U%s}%%{+u}%s", c, v)
	return s
}

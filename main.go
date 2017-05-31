package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

var arithmetic = map[string]bool{
	"+": true,
	"-": true,
	"*": true,
	"/": true,
}

func skip(p string) string {
	if len(p) == 0 {
		return ""
	}
	for ; p[0] == ' '; p = p[1:] {
	}
	return p
}

func readToken(p string) (string, string) {
	p = skip(p)
	i := 0
	for ; i < len(p) && p[i] != ' '; i++ {
	}
	return p[0:i], p[i:len(p)]
}

func eval(p string) (int, string) {
	t, p := readToken(p)

	// Number
	if '0' <= t[0] && t[0] <= '9' {
		i, _ := strconv.Atoi(t)
		return i, p
	}

	// Arithmetic operators
	if _, ok := arithmetic[t]; ok {
		x, p := eval(p)
		y, p := eval(p)
		switch t {
		case "+":
			return x + y, p
		case "-":
			return x - y, p
		case "*":
			return x * y, p
		case "/":
			return x / y, p
		}
	}

	panic("unsupported operator: " + t)
}

func main() {
	x, _ := eval(strings.Join(os.Args[1:], " "))
	fmt.Println(x)
}

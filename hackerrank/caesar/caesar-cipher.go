package main

import (
	"fmt"
	"unicode"
)

var (
	aUpper = byte('A')
	aLower = byte('a')
)
var sz, sw int
var s string

func cipher() string {
	output := make([]byte, sz, sz)
	var l byte
	for _, v := range s {
		if unicode.IsUpper(v) {
			l = move(v, aUpper)
		} else if unicode.IsLower(v) {
			l = move(v, aLower)
		} else {
			l += byte(v)
		}
		output = append(output, l)
	}
	return string(output)
}

func move(v rune, base byte) byte {
	l := (byte(v) - base) + byte(sw)
	return (l % 26) + base
}

func main() {
	fmt.Scanf("%d\n%s\n%d", &sz, &s, &sw)
	fmt.Println(cipher())
}

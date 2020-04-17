package main

import (
	"fmt"
)

func main() {
	var s string
	a := byte('A')
	z := byte('Z')
	fmt.Scanf("%s", &s)
	if len(s) == 0 {
		fmt.Println("0")
		return
	}
	count := 1
	for i := 0; i < len(s); i++ {
		if s[i] >= a && s[i] <= z {
			count++
		}
	}
	fmt.Println(count)
}

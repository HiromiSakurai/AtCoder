package main

import (
	"fmt"
	"strings"
)

func main() {
	var S string
	var T string

	fmt.Scan(&S)

	strs := []string{"dream", "dreamer", "erase", "eraser"}

	// 先頭からチェックしていくと、dream/dreamerの違いを判別できない
	// 末尾からチェックすることによって上記の問題を解決するために、文字列を反転させる
	rS := reverse(S)
	for i, s := range strs {
		strs[i] = reverse(s)
	}

	can := true
	next := true

	for next {
		for i, s := range strs {
			tmp := T + s

			if strings.HasPrefix(rS, tmp) {
				T = tmp

				if len(T) == len(rS) {
					next = false
				}
				break
			}

			if i == len(strs)-1 {
				can = false
				next = false
			}
		}
	}

	if can {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}

func reverse(input string) string {
	runes := []rune(input)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

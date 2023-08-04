package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var N int
	fmt.Scanf("%d", &N)

	sc := bufio.NewScanner(os.Stdin)

	set := make(map[string]struct{})

	for i := 0; i < N; i++ {
		sc.Scan()
		input := sc.Text()
		set[input] = struct{}{}
	}

	fmt.Println(len(set))
}

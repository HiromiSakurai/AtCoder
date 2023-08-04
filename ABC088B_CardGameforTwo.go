package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	var N int
	fmt.Scanf("%d", &N)

	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	inputs := strings.Split(sc.Text(), " ")

	var arr []int
	for _, input := range inputs {
		p, _ := strconv.Atoi(input)
		arr = append(arr, p)
	}

	descending(arr)

	alice := 0
	bob := 0

	for i, e := range arr {
		if even(i) {
			alice += e
		} else {
			bob += e
		}
	}

	fmt.Println(alice - bob)
}

func even(n int) bool {
	return n%2 == 0
}

func descending(sl []int) {
	sort.Sort(sort.Reverse(sort.IntSlice(sl)))
}

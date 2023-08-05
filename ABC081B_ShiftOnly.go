package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var N int
	fmt.Scanf("%d", &N)

	sc := bufio.NewScanner(os.Stdin)

	nums := []int{}

	sc.Scan()
	inputs := strings.Split(sc.Text(), " ")

	for _, v := range inputs {
		num, _ := strconv.Atoi(v)
		nums = append(nums, num)
	}

	count := 0
	next := true

	for next {
		for i, v := range nums {
			if v%2 != 0 {
				next = false
				break
			}

			nums[i] = v / 2
			if i == len(nums)-1 {
				count++
			}
		}
	}

	fmt.Println(count)
}

package main

import "fmt"

func main() {
	var s1, s2, s3 int
	fmt.Scanf("%1d%1d%1d", &s1, &s2, &s3)

	sl := []int{s1, s2, s3}

	count := 0

	for _, v := range sl {
		if v == 1 {
			count++
		}
	}

	fmt.Println(count)
}

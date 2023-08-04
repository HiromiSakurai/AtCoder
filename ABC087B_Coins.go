package main

import "fmt"

func main() {
	var A int // 500yen
	var B int // 100yen
	var C int // 50yen
	var X int // total amount

	fmt.Scanf("%d", &A)
	fmt.Scanf("%d", &B)
	fmt.Scanf("%d", &C)
	fmt.Scanf("%d", &X)

	count := 0

	for i := 0; i <= A; i++ {
		for j := 0; j <= B; j++ {
			for k := 0; k <= C; k++ {
				if 500*i+100*j+50*k == X {
					count++
				}
			}
		}
	}

	fmt.Println(count)
}

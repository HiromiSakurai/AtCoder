package main

import "fmt"

func main() {
	var N int
	var A int
	var B int

	fmt.Scanf("%d", &N)
	fmt.Scanf("%d", &A)
	fmt.Scanf("%d", &B)

	result := 0

	for i := 1; i <= N; i++ {
		num := i
		sum := sumOfDigits(num)

		if A <= sum && sum <= B {
			result += i
		}
	}
	fmt.Println(result)
}

// 数値の各桁の和を計算する関数
func sumOfDigits(num int) int {
	sum := 0
	// 各桁の和を計算
	for num != 0 {
		// numを10で割った余りが各桁の数値
		digit := num % 10
		sum += digit
		// numを10で割って次の桁へ
		num /= 10
	}
	return sum
}

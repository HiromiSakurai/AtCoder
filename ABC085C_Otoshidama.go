package main

import "fmt"

func main() {
	var N, Y int
	fmt.Scanf("%d %d", &N, &Y)

	var _10000, _5000, _1000 int = -1, -1, -1

OuterLoop:
	for i := 0; i <= N; i++ {
		for j := 0; j <= N-i; j++ {

			// 1000円札の枚数を計算
			k := N - i - j
			if 10000*i+5000*j+1000*k == Y {
				_10000, _5000, _1000 = i, j, k

				// ラベル付きbreak文で多重ループを抜ける
				break OuterLoop
			}
		}
	}

	fmt.Printf("%d %d %d\n", _10000, _5000, _1000)
}

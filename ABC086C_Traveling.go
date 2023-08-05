package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	var N int
	fmt.Scanf("%d", &N)

	sc := bufio.NewScanner(os.Stdin)

	ts := []int{0}
	xs := []int{0}
	ys := []int{0}

	for i := 0; i < N; i++ {
		sc.Scan()
		inputs := strings.Split(sc.Text(), " ")
		t, _ := strconv.Atoi(inputs[0])
		x, _ := strconv.Atoi(inputs[1])
		y, _ := strconv.Atoi(inputs[2])
		ts = append(ts, t)
		xs = append(xs, x)
		ys = append(ys, y)
	}

	can := true

	// 時間差dt, 距離dist(マンハッタン距離)とした時、以下の条件を満たす場合に有効な移動と考えられる
	// 1. dt >= dist
	// 2. dtとdistの偶奇が一致
	// ※ マンハッタン距離とは、x軸方向とy軸方向の差の絶対値の和
	for i := 0; i < N; i++ {
		dt := ts[i+1] - ts[i]

		distf := math.Abs(float64(xs[i+1]-xs[i])) + math.Abs(float64(ys[i+1]-ys[i]))
		dist := int(distf)

		// 条件 1
		if dt < dist {
			can = false
			break
		}
		// 条件 2
		if dist%2 != dt%2 {
			can = false
			break
		}
	}

	if can {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}

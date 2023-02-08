package main

import (
	"fmt"
	"math"
)

func main() {
	var N int
	fmt.Scan(&N)
	h := make([]float64, N)
	for i:=0; i<N; i++ {
		fmt.Scan(&h[i])
	}

	//配列dpを定義（配列全体を無限大を表す値に初期化）
	dp := make([]float64, N)
	for i:=0; i<N; i++ {
		dp[i] = math.MaxFloat64
	}

	//初期条件
	dp[0] = 0

	//ループ
	for i:=1; i<N; i++ {
		if i==1 {
			dp[i] = math.Abs(h[i]-h[i-1])
		} else {
			x := dp[i-1] + math.Abs(h[i]-h[i-1])
			y := dp[i-2] + math.Abs(h[i]-h[i-2])
			dp[i] = math.Min(x, y)
		}
	}
	fmt.Println(dp[N-1])
}

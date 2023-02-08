package main

import "fmt"

var memo [50]int64

func fibo(N int) int64 {
	//ベースケース
	if N==0 {
		return 0
	} else if N==1 {
		return 1
	}

	//メモをチェック（既に計算済みなら答えをリターン）
	if memo[N] != 0 {
		return memo[N]
	}

	//答えをメモ化しながら、再帰呼び出し
	memo[N] = fibo(N-1) + fibo(N-2)
	return memo[N]
}

func main() {
	fibo(49)

	for N:=2; N<50; N++ {
		fmt.Printf("%d項目%v\n", N, memo[N])
	}
}

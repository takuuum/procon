package main

import "fmt"

func main() {
	// 変数を定義
	var tanak, shohizei float64
	shohizei = 1.1

	// 入力
	fmt.Scan(&tanak)

	// 計算
	gokei := tanak * shohizei

	// 出力
	fmt.Println(int(gokei))
}

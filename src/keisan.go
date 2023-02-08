package main

import "fmt"

func main() {

	// 入力
	var 縦, 横 int
	fmt.Scan(&縦, &横)

	if 縦<0 || 横<0 {
		fmt.Println("正の数を入力してください")
		return
	}

	// 計算
	var 面積 int
	// 15 = 3 * 5
	面積 = 縦 * 横

	// 出力
	fmt.Println("面積：",面積)

}

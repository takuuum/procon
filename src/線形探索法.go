package main

import (
	"fmt"
)

func main() {
	// 入力を受け取る
	var N, v int
	fmt.Scan(&N,&v)
	a := make([]int, N)
	for i:=0; i<N; i++ {
		fmt.Scan(&a[i])
	}

	// 線形探索
	exist := false
	found_id := -1
	for i:=0; i<N; i++ {
		if a[i] == v {
			exist = true
			found_id = i
		}
	}

	// 結果出力
	if exist {
		fmt.Println("Yes: ", found_id)
	} else {
		fmt.Println("No")
	}
}

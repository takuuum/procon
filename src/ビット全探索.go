package main

import "fmt"

func main() {
	// 入力を受け取る
	var N, W int
	fmt.Scan(&N,&W)
	a := make([]int, N)
	for i:=0; i<N; i++ {
		fmt.Scan(&a[i])
	}

	// bitは2^N通りの部分集合全体を動く
	exist := false
	//0 = 000
	//1 = 001
	//2 = 010
	//3 = 011
	//4 = 100
	//5 = 101
	//6 = 110
	//7 = 111
	for bit:=0; bit<(1<<N); bit++ {
		sum:=0
		for i:=0; i<N; i++ {
			//i番目の要素a[i]が部分集合に含まれているかどうか
			if (bit & (1<<i)) {
				sum += a[i]
			}
		}

		// sumがWに一致するかどうか
		if sum == W {
			exist = true
		}
	}

	if exist {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}

package main

import "fmt"

func main() {
	var N int
	fmt.Scan(&N)
	S := make([]int, N)
	T := make([]int, N)

	for i:=0; i<N; i++ {
		fmt.Scan(&S[i])
	}
	for i:=0; i<N; i++ {
		fmt.Scan(&T[i])
	}
	t := T[0]
	fmt.Println(t)

	for i:=1; i<N; i++ {
		t += S[i-1]
		S[i-1] = t
		if S[i-1] <= T[i] {
			fmt.Println(S[i-1])
		} else {
			fmt.Println(T[i])
			t = T[i]
		}
	}
}
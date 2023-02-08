package main

import "fmt"

func main() {
	var N int
	fmt.Scan(&N)

	name := make([][]string, N)

	for i:=0; i<N; i++ {
		name[i] = make([]string, 2)
		fmt.Scan(&name[i][0])
		fmt.Scan(&name[i][1])
	}

	count := 0

	for i:=0; i<N; i++ {
		for j:=0; j<N; j++ {
			if name[i][0] == name[j][0] && name[i][1] == name[j][1] {
				count++
			}
		}
	}
	if count>N {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}

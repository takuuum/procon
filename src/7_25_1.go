package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	numMap := make(map[int]string)
	for i:=0; i<n; i++ {
		var q int
		fmt.Scan(&q)
		if q==1 {
			var s string
			fmt.Scan(&s)
			var a int
			fmt.Scan(&a)
			numMap[a] = s
		} else {
			var a int
			fmt.Scan(&a)
			if val, ok := numMap[a]; ok {
				fmt.Println(val)
			} else {
				fmt.Println("NULL")
			}
		}
	}
}

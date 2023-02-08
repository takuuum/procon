package main

import "fmt"

func main() {
	var q int
	fmt.Scan(&q)
	m := map[int]string{}

	for i:=0; i<q; i++ {
		var s int
		var u int
		fmt.Scan(&s)
		if s == 2 {
			fmt.Scan(&u)
			v, ok := m[u]
			if ok {
				fmt.Println(v)
			} else {
				fmt.Println("NULL")
			}
		} else {
			var t string
			fmt.Scan(&t, &u)
			m[u] = t
		}
	}
}

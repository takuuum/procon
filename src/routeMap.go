package main

import "fmt"

func main() {
	var n,m int
	fmt.Scan(&n,&m)
	s := make(map[string]bool, n)
	ss := make([]string, n)
	var key string
	for i:=0; i<n; i++ {
		fmt.Scan(&ss[i])
		s[ss[i]] = false
	}
	for i:=0; i<m; i++ {
		fmt.Scan(&key)
		_, ok := s[key]
		if ok {
			s[key] = true
		}
	}

	for i:=0; i<n; i++ {
		if s[ss[i]] {
			fmt.Println("Yes")
		} else {
			fmt.Println("No")
		}
	}
}

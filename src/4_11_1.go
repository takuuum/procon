package main

import (
	"fmt"
	"strings"
)

func main() {
	var s string
	fmt.Scan(&s)
	ss := strings.Split(s, "")
	tt := make([]string, 0)
	c := 0
	for i:=0; i<len(ss); i++ {
		count := 0
		for j:=0; j<len(tt); j++ {
			if ss[i] == tt[j] {
				count++
			}
		}
		if count == 0 {
			tt = append(tt, ss[i])
			c++
		}
	}
	fmt.Println(c)
}

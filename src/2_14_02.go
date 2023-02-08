package main

import (
	"fmt"
	"strings"
)

func main() {
	var s string
	k := 0
	fmt.Scan(&s)
	ss := strings.Split(s,"")
	for i:=1; i<len(ss); i++ {
		if k == 0 && ss[k] == ss[i] {
			k = i
		} else if 0 < k && ss[i-k] == ss[i] {
			continue
		} else {
			k = 0
		}
	}
	if k == 0 {
		k = len(ss)
	}
	fmt.Println(k)
}

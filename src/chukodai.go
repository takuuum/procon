package main

import (
	"fmt"
	"strings"
)

func main() {
	var s string
	var a,b int
	fmt.Scan(&s,&a,&b)

	ss := strings.Split(s,"")
	aa := ss[a-1]
	bb := ss[b-1]
	ss[a-1] = bb
	ss[b-1] = aa

	sss := ""
	for i:=0; i<len(ss); i++ {
		sss += ss[i]
	}
	fmt.Println(sss)
}

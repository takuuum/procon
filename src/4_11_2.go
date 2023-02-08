package main

import (
	"fmt"
	"strings"
)

func main() {
	var s string
	var n int
	fmt.Scan(&s, &n)
	c := make([]string, n)
	aa := make([]string, 0)
	for i:=0; i<n; i++ {
		fmt.Scan(&c[i])
		if len(s) == len(c[i]) {
			var d string
			for j:=0; j<len(s); j++ {
				if j==0 {
					d = c[i]
				}
				if s == d {
					aa = append(aa, c[i])
					break
				}
				cc := strings.Split(d, "")
				d = ""
				for k:=0; k<len(cc)-1; k++ {
					d += cc[k+1]
				}
				d += cc[0]
			}
		}
	}
	for i:=0; i<len(aa); i++ {
		fmt.Println(aa[i])
	}
}

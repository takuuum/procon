package main

import "fmt"

func main() {
	p := make([]int, 26)
	s := make([]string, 26)

	for i:=0;i<26;i++ {
		fmt.Scan(&p[i])
		s[i] = toChar(p[i])
		fmt.Print(s[i])
	}
}

func toChar(i int) string {
	return string('a'-1 + i)
}

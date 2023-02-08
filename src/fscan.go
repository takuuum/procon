package main

import (
	"fmt"
	"strings"
)

var source = "123 1.234 1.0e4 test"

func main() {
	reader := strings.NewReader(source)
	var i int
	var f,g float64
	var s string
	a,_ := fmt.Println(reader)
	fmt.Println(a)
	fmt.Fscan(reader, &i, &f, &g, &s)
	a,_ = fmt.Printf("i=%v f=%v g=%v s=%v\n", i, f, g, s)
	fmt.Println(a)
	a,_ = fmt.Println(1)
	fmt.Println(a)
}

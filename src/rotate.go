package main

import "fmt"

func main() {
	var a int
	fmt.Scan(&a)
	b := a/100
	c := (a-b*100)/10
	d := a-b*100-c*10
	// bcd
	// cdb
	e := c*100 + d*10 + b
	// dbc
	f := d*100 + b*10 + c

	fmt.Println(a+e+f)
}

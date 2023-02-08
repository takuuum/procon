package main

import "fmt"

var a,b int

func main() {
	var l,q int
	fmt.Scan(&l,&q)

	cx := make([][]int, q)
	wood := make([]bool, l+1)
	wood[0] = true
	wood[l] = true

	for i:=0; i<q; i++ {
		cx[i] = make([]int, 2)
		fmt.Scan(&cx[i][0],&cx[i][1])
		if cx[i][0] == 1 {
			wood[cx[i][1]] = true
		} else if cx[i][0] == 2 {
			go search1(wood, cx, i)
			search2(wood, cx, i)
			fmt.Println(a+b)
		}
	}
}

func search1(wood []bool, cx [][]int, i int)  {
	j := cx[i][1]
	count := 0

	for ;; {
		count++
		if wood[j-1] == false {
			j--
			continue
		} else {
			a = count
		}
	}
}

func search2(wood []bool, cx [][]int, i int) {
	j := cx[i][1]
	count := 0

	for ;; {
		count++
		if wood[j+1] == false {
			j++
			continue
		} else {
			b = count
		}
	}
}
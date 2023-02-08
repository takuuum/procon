package main

import (
	"fmt"
	"strings"
)

func main() {
	var r, c int
	fmt.Scan(&r, &c)
	var si, sj int
	cc := make([][]string, r)
	for i:=0; i<r; i++ {
		cc[i] = make([]string, c)
		var a string
		fmt.Scan(&a)
		fmt.Println("c")
		for j:=0; j<c; j++ {
			cc[i][j] = strings.Split(a, "")[c]
			if cc[i][j] == "S" {
				si = i
				sj = j
			}
		}
	}
	fmt.Println("a")
	// earlist := make([]int, 0)
	r--
	c--
	sum := do(r, c, si, sj, 0, cc)
	fmt.Println(sum)
}

func do(r, c, ci, cj, sum int, cc [][]string) int {
	fmt.Println(ci, ",", ci, ",", sum, ",", cc[ci][cj])
	if cc[ci][cj] == "T" {
		return sum
	}
	cc[ci][cj] = "."
	var summ int
	if ci!=r && cc[ci+1][cj] != "." {
		summ = do(r, c, ci+1, cj, sum+1, cc)
		if summ<sum {
			sum = summ
		}
	}
	if cj!=c && cc[ci][ci] != "." {
		summ = do(r, c, ci, cj+1, sum+1, cc)
		if summ<sum {
			sum = summ
		}
	}
	if cc[ci-1][cj] != "." {
		summ = do(r, c, ci-1, cj, sum+1, cc)
		if summ<sum {
			sum = summ
		}
	}
	if cc[ci][cj-1] != "." {
		summ = do(r, c, ci-1, cj, sum+1, cc)
		if summ<sum {
			sum = summ
		}
	}
	return sum
}

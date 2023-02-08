package main

import (
	"fmt"
	"reflect"
	"sort"
)

func main() {
	var N int
	fmt.Scan(&N)
	s := make([][]string, N)
	t := make([]int, N*N)
	u := make([]int, N*N)

	for i:=0;i<N*2;i++ {
		s[i] = make([]string, N*2)
		for j:=0;j<N;j++ {
			fmt.Scan(&s[i][j])
		}
	}

	for i:=0;i<N*2;i++ {
		for j:=0;j<N;j++ {
			count := 0
			if s[i][j] == "#" {
				if s[j][j-1] == "#" {
					count++
				}
				if s[i][j+1] == "#" {
					count++
				}
				if s[i-1][j] == "#" {
					count++
				}
				if s[i+1][j] == "#" {
					count++
				}
				if i<N {
					t[j*(i+1)] = count
				} else {
					u[j*(i+1)] = count
				}
			} else {
				if i<N {
					t[j*(i+1)] = 0
				} else {
					u[j*(i+1)] = 0
				}
			}
		}
	}

	sort.Slice(t, func(i, j int) bool {
		return t[i] < t[j]
	})
	sort.Slice(u, func(i, j int) bool {
		return u[i] < u[j]
	})

	if reflect.DeepEqual(t,u) == true {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}

package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	t := make([]int64, n)
	k := make([]int, n)
	a := make([][]int, n)
	var count int64
	count = 0
	for i:=0; i<n; i++ {
		fmt.Scan(&t[i], &k[i])
		a[i] = make([]int, k[i]+1)
		if k[i] == 0 {
			a[i][0] = 0
			continue
		}
		for j:=0; j<k[i]; j++ {
			fmt.Scan(&a[i][j])
			if i==n-1 {
				count += t[j]
			}
		}
	}
	fmt.Println(a, a[n-1][0], a[n-1], count, t)
	for i:=0; i<len(a[n-1]); i++ {
		count = aaa(a, a[n-1][i], a[n-1], count, t)
	}
	count += t[n-1]
	fmt.Println(count)
}

func aaa(a [][]int, i int, aa []int, count int64, t []int64) int64{
	if a[i][0] == 0 {
		return count
	}
	for j:=0; j<len(a[i]); j++ {
		if contains(aa, a[i][j]) {
			continue
		} else {
			count += t[i]
		}
		return aaa(a, a[i][j], aa, count, t)
	}
	return count
}

func contains(s []int, e int) bool {
	for _, v := range s {
		if e == v {
			return true
		}
	}
	return false
}

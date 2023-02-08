package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	x := make([]int, 12)
	y := make([]int, 12)
	for i:=0; i<12; i++ {
		fmt.Scan(&x[i])
		y[i] = x[i]
	}

	a := make([]float64, 3)
	b := make([]float64, 3)

	x[2] -= x[0]
	x[3] -= x[1]
	x[4] -= x[0]
	x[5] -= x[1]

	x[8] -= x[6]
	x[9] -= x[7]
	x[10] -= x[6]
	x[11] -= x[7]

	c := x[2]*y[4]+x[3]*x[5]
	d := math.Sqrt(float64(x[2]*x[2]+x[3]*x[3]))*math.Sqrt(float64(x[4]*x[4]+x[5]*x[5]))
	a[0] = float64(c)/d

	c = x[8]*y[10]+x[9]*x[11]
	d = math.Sqrt(float64(x[8]*x[8]+x[9]*x[9]))*math.Sqrt(float64(x[10]*x[10]+x[11]*x[11]))
	b[0] = float64(c)/d

	y[0] -= y[2]
	y[1] -= y[3]
	y[4] -= y[2]
	y[5] -= y[3]

	y[6] -= y[8]
	y[7] -= y[9]
	y[10] -= y[8]
	y[11] -= y[9]

	c = x[0]*y[4]+x[1]*x[5]
	d = math.Sqrt(float64(x[0]*x[0]+x[1]*x[1]))*math.Sqrt(float64(x[4]*x[4]+x[5]*x[5]))
	a[1] = float64(c)/d

	c = x[6]*y[10]+x[7]*x[11]
	d = math.Sqrt(float64(x[6]*x[6]+x[7]*x[7]))*math.Sqrt(float64(x[10]*x[10]+x[11]*x[11]))
	b[1] = float64(c)/d

	a[2] = math.Abs(a[0]-a[1])
	b[2] = math.Abs(b[0]-b[1])
	sort.Slice(a, func(i, j int) bool { return a[i] < a[j] })
	sort.Slice(b, func(i, j int) bool { return b[i] < b[j] })

	for i:=0; i<3; i++ {
		if a[i]!=b[i] {
			fmt.Println("No")
			return
		}
	}
	fmt.Println("Yes")
}

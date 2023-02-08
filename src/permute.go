package main

import "fmt"

// リストからidx番目の要素を抜かしたものを取得
func delete(idx int, L []interface{}) (result []interface{}) {
	result = append(result, L[:idx]...)
	result = append(result, L[idx+1:]...)
	return
}

// sliceの全組み合わせを返す
func Permute(L []interface{}) (result [][]interface{}) {

	var inner func(LL []interface{})

	inner = func(LL []interface{}) {
		if len(LL) == 0 {
			result = append(result, []interface{}{}) //
		}
		for idx, r := range L {
			for _, t := range Permute(delete(idx, L)) {
				result = append(result, append([]interface{}{r}, t...))
			}
		}
	}
	inner(L)
	return

}

func main() {
	for _, l := range Permute([]interface{}{'a', 'b', 'c'}) {
		// lに全ての組み合わせがrune型で返ってくるから、それをstring型にして出力
		result := make([]rune, len(l))
		for i, e := range l {
			result[i] = e.(rune)
		}
		fmt.Println(string(result))
	}
}

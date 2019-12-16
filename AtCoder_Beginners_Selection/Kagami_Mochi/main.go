package main

import (
	"fmt"
)

func main() {
	n := 0
	x := 0
	m := make(map[int]struct{})
	var a []int
	var new []int

	//入力
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		fmt.Scan(&x)
		a = append(a, x)
	}
	//処理
	for _, element := range a {
		if _, ok := m[element]; !ok {
			m[element] = struct{}{}
			new = append(new, element)
		}
	}
	//出力
	fmt.Print(len(new))

}

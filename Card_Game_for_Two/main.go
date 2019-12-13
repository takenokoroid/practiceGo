package main

import (
	"fmt"
	"sort"
)

func main() {
	var n int
	var x int
	var a []int
	Alise := 0
	Bob := 0

	fmt.Scanf("%d", &n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &x)
		a = append(a, x)
	}

	sort.Sort(sort.IntSlice(a))
	for i := 0; i < n; i++ {
		Alise = Alise + a[len(a)-1]
		a = unset(a, len(a)-1)
		i++
		if i == n {
			break
		}
		Bob = Bob + a[len(a)-1]
		a = unset(a, len(a)-1)
	}
	fmt.Printf("%d", Alise-Bob)
}

func unset(s []int, i int) []int {
	if i >= len(s) {
		return s
	}
	return append(s[:i], s[i+1:]...)
}

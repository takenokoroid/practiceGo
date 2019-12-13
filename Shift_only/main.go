package main

import "fmt"

func main() {
	var N int
	var a int
	count := 0
	min := 10000
	fmt.Scanf("%d", &N)
	for i := 0; i < N; i++ {
		fmt.Scanf("%d", &a)
		for count = 0; ; count++ {
			if a%2 == 0 {
				a = a / 2
			} else {
				break
			}
		}
		if min > count {
			min = count
		}
	}
	fmt.Printf("%d", min)
}

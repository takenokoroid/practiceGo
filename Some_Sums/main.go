package main

import "fmt"

func main() {
	var n, a, b int
	ans := 0
	fmt.Scanf("%d %d %d", &n, &a, &b)
	for i := 1; i <= n; i++ {
		sum := 0
		dig := 0
		j := i
		for j > 0 {
			dig = j % 10
			sum = sum + dig
			j = j / 10
		}
		if a <= sum && b >= sum {
			ans = ans + i
		}
	}
	fmt.Printf("%d", ans)
}

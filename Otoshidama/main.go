package main

import (
	"fmt"
)

func main() {
	n := 0
	sum := 0
	//入力
	fmt.Scan(&n, &sum)
	//処理
	for i := 0; i <= n; i++ {
		for j := 0; j <= (n - i); j++ {
			k := n - (i + j)
			total := i*10000 + j*5000 + k*1000
			if total == sum {
				fmt.Println(i, j, k)
				return
			}
		}
	}
	//解なし出力
	fmt.Println(-1, -1, -1)
}

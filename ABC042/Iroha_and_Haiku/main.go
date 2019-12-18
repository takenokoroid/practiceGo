package main

import "fmt"

func main() {
	var a int
	count5, count7 := 0, 0
	for i := 0; i < 3; i++ {
		//input
		fmt.Scan(&a)

		//process
		if a == 5 {
			count5++
		} else if a == 7 {
			count7++
		}
	}
	if count5 == 2 && count7 == 1 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}

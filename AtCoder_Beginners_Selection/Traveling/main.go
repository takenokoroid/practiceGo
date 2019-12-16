package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	var t, x, y, t2, x2, y2 float64
	//input
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		var distance float64
		var dt float64
		fmt.Scan(&t, &x, &y)

		//process
		dt = t - t2
		distance = math.Abs(x-x2) + math.Abs(y-y2)
		if dt < distance || int(dt)%2 != int(distance)%2 {
			fmt.Println("No")
			return
		}
		t2 = t
		x2 = x
		y2 = y
	}
	fmt.Println("Yes")
}

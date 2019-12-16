package main

import (
	"fmt"
	"strings"
)

func main() {
	var s string
	//input
	fmt.Scan(&s)
	//process
	s = strings.Replace(s, "dream", "D", -1)
	s = strings.Replace(s, "erase", "E", -1)
	s = strings.Replace(s, "Der", "", -1)
	s = strings.Replace(s, "Er", "", -1)
	s = strings.Replace(s, "D", "", -1)
	s = strings.Replace(s, "E", "", -1)

	if s == "" {
		fmt.Print("YES")
		return
	} else {
		fmt.Print("NO")
	}
}

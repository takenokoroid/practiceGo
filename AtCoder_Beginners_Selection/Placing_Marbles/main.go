package main

import (
	"fmt"
	"strings"
)

func main() {
	var s string
	fmt.Scanf("%s", &s)
	fmt.Print(strings.Count(s, "1"))
}

package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b float64
	fmt.Scan(&a)
	fmt.Scan(&b)
	r := math.Pow(a, 2) + math.Pow(b, 2)
	fmt.Println(math.Sqrt(r))
}

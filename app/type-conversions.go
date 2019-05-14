package main

import (
	"fmt"
	"math"
)

func main() {
	var x, y int = 3, 4
	// 型変換
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)
	// 3 4 5
	fmt.Println(x, y, z)
}
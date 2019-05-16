package main

import (
	"fmt"
	"math"
)

func pow(x, n, lim float64) float64 {
	// xのn乗がlimより小さいか
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		// 	27 >= 20
		fmt.Printf("%g >= %g\n", v, lim)
	}
	// can't use v here, though
	return lim
}

func main() {
	// 9 20
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)
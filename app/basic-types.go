package main

import (
	"fmt"
	"math/cmplx"
)

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

func main() {
	// Type: bool Value: false
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	// Type: uint64 Value: 18446744073709551615
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	// Type: complex128 Value: (2+3i)
	fmt.Printf("Type: %T Value: %v\n", z, z)
}
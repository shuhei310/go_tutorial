package main

import (
	"fmt"
	"math"
)
// エクスポートされていない名前は、外部のパッケージからアクセスすることはできません。
func main() {
	// prog.go:9:14: cannot refer to unexported name math.pi
	// prog.go:9:14: undefined: math.pi
	fmt.Println(math.pi)
}
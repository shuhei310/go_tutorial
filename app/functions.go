package main

import "fmt"

func add(x int, y int) int {
	return x + y
}

func main() {
	// add関数で足し算
	// 55
	fmt.Println(add(42, 13))
}
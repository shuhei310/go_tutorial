package main

import "fmt"

// naked returnステートメント
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func main() {
	// 7 10
	fmt.Println(split(17))
}
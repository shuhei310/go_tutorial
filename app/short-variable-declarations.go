package main

import "fmt"

func main() {
	var i, j int = 1, 2
	// var宣言の代わりに、短い := の代入文を使い、暗黙的な型宣言ができる
	k := 3
	c, python, java := true, false, "no!"

	// 1 2 3 true false no!
	fmt.Println(i, j, k, c, python, java)
}
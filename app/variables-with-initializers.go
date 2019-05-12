package main

import "fmt"

// var宣言では、変数毎に初期化子を与えることができる
var i, j int = 1, 2

func main() {
	// 初期化子が与えられている場合、型を省略できる
	var c, python, java = true, false, "no!"
	// 1 2 true false no!
	fmt.Println(i, j, c, python, java)
}
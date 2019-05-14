package main

import "fmt"

func main() {
	// 明示的な型を指定せずに変数を宣言する場合、変数の型は右側の変数から型推論される
	v := 42 // change me!
	// v is of type int
	fmt.Printf("v is of type %T\n", v)
}
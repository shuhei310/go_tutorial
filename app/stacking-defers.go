package main

import "fmt"

func main() {
	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		// deferへ渡した順に実行される
		defer fmt.Println(i)
	}

	fmt.Println("done")
}
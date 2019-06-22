package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{1, 2}
	p := &v
	// (*p).Xの代わり
	p.X = 1e9
	// {1000000000 2}
	fmt.Println(v)
}
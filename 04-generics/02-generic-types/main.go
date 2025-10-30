package main

import "fmt"

type List[T any] struct {
	next *List[T]
	val T
}

func main() {
	x := List[int]{nil, 10}
	fmt	.Println(x)

	y := List[string]{nil, "hello"}
	fmt.Println(y)
}
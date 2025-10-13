package main

import "fmt"

func fibonacci() func(int) int {
	return func(x int) int {
		a, b := 0, 1
		for range x {
			a, b = b, a+b
		}

		return a
	}
}

func main() {
	f := fibonacci()

	for i := 0; i < 10; i++ {
		fmt.Println(f(i))
	}
}

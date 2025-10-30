package main

import "fmt"

type I interface {
	M()
}

func main() {
	var i I

	// IStruct を代入すれば動きはする
	// var i I = &IStruct{}
	describe(i)
	i.M()
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}

type IStruct struct{}

func (i *IStruct) M() {
	fmt.Println("IStruct M method called")
}

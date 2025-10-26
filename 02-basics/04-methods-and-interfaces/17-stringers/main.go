package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

// NOTE: fmt などのパッケージは引数に渡された型が Stringer インターフェースをチェックしており、
// 定義されていた場合は自動的に String メソッドを呼び出してその返り値を出力する。
func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

func main() {
	a := Person{"Arthur Dent", 42}
	z := Person{"Zaphod Beeblebrox", 9001}
	fmt.Println(a, z)
}
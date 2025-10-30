package main

import (
	"fmt"
	"image"
)

func main() {
	m := image.NewRGBA(image.Rect(255, 255, 255, 255))
	fmt.Println(m.Bounds())
	fmt.Println(m.At(0, 0).RGBA())
}
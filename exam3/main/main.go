package main

import (
	"correction/exam3"
	"fmt"
)

func main() {
	exam3.Printdigits()
	exam3.Displayalpham()
	exam3.Displayalrevm()
	fmt.Println(exam3.Rot14("Hello !"))
	a := []int{23, 123, 1, 11, 55, 93}
	fmt.Println(exam3.Max(a))
	fmt.Println(exam3.Compare("Hello!", "Hello!"))
	fmt.Println(exam3.Compare("Salut!", "lut!"))
	fmt.Println(exam3.Compare("Ola!", "Ol"))
}

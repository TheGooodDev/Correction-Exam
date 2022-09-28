package main

import (
	"correction/exam2"
	"fmt"

	"github.com/01-edu/z01"
)

func main() {
	exam2.Displaya()
	exam2.Displayz()
	exam2.Hello()
	z01.PrintRune((exam2.FirstRune("hello")))
	z01.PrintRune('\n')
	z01.PrintRune((exam2.Lastrune("hello")))
	z01.PrintRune('\n')
	z01.PrintRune((exam2.Nrune("hello", 2)))
	z01.PrintRune('\n')
	fmt.Println(exam2.StrRev("hello"))
	exam2.PrintStr("hello printstr")
	fmt.Println(exam2.StrLen("Hello World!"))
	fmt.Println(exam2.Rot14("Hello! How are You?"))
	fmt.Println(exam2.Atoi("51535235"))
}

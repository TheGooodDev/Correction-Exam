package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args
	if len(args) != 4 || len(args[2]) != 1 || len(args[3]) != 1 {
		return
	}
	trune := []rune(args[1])
	for i := 0; i < len(args[1]); i++ {
		if string(trune[i]) == args[2] {
			trune[i] = rune(args[3][0])
		}
	}
	fmt.Println(string(trune))
}

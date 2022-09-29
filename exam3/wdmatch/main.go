package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args
	if len(args) != 3 {
		return
	}
	count := 0
	for _, c := range args[2] {
		if c == rune(args[1][count]) {
			count++
		}
		if count == len(args[1]) {
			fmt.Println(args[1])
			return
		}
	}

}

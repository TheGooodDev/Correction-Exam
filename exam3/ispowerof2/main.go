package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args
	if len(args) != 2 {
		return
	}
	res := 0
	for i, c := range args[1] {
		res += int(c - 48)
		if i != len(args[1])-1 {
			res *= 10
		}
	}
	for power := 2; power <= res; power *= 2 {
		if power == res {
			fmt.Println("true")
			return
		}
	}
	fmt.Println("false")
}

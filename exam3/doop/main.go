package main

import (
	"correction/utils"
	"fmt"
	"os"
)

func main() {
	args := os.Args
	if len(args) != 4 {
		return
	}
	nb1, err := utils.Atoi(args[1])
	nb2, err2 := utils.Atoi(args[3])
	op := args[2]
	if err != nil || err2 != nil {
		return
	}
	switch op {
	case "+":
		fmt.Println(nb1 + nb2)
	case "-":
		fmt.Println(nb1 - nb2)
	case "*":
		fmt.Println(nb1 * nb2)
	case "/":
		if nb2 == 0 {
			fmt.Println("No division by 0")
			return
		}
		fmt.Println(nb1 / nb2)
	case "%":
		if nb2 == 0 {
			fmt.Println("No modulo by 0")
			return
		}
		fmt.Println(nb1 % nb2)
	}
}

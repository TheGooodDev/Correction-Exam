package main

import (
	"fmt"
	"os"
)

func main() {
	arg := os.Args
	fmt.Println(len(arg) - 1)
}

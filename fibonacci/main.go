
package main

import (
	"os"
	"fmt"
	"strconv"
)

func main () {
	if len(os.Args) < 2 {
		fmt.Println("arguments required: target fibonacci index")
		return
	}

	input, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}

	result := nextFibonacci(input)

	fmt.Println("target index  =", result)
}

func nextFibonacci(idx int) int {
	if idx <= 1 {
		return 1
	}

	return nextFibonacci(idx - 1) + nextFibonacci(idx - 2)
}


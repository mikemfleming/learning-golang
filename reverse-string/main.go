package main

import "fmt"
import "os"

func main() {
	var input = os.Args[1]
	var output string
	for i := len(input) - 1; i > -1; i-- {
		output += string(input[i])
	}
	fmt.Println(output)
}


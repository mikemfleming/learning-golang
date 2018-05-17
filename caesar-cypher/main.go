
package main

import "flag"
import "fmt"
import "strings"

func main() {
    message := flag.String("message", "hello_world", "the message to encode or decode")
    offset := flag.Int("offset", 1, "the offset to encode the message with")
    command := flag.String("command", "encode", "the operation to perform")
    flag.Parse()

    if *command == "encode" {
    	fmt.Println("encoded:", cypher(*message, *offset))
    }

    if *command == "decode" {
    	fmt.Println("decoded:", cypher(*message, -1 * *offset))
    }
}

func cypher(msg string, offset int) string {
	var alphabet = "abcdefghijklmnopqrstuvwxyz_"
	var encodedMsg string

	for i := 0; i < len(msg); i++ {
		oldChar := string(msg[i]) // strings are arrays of bytes
		newIdx := (strings.Index(alphabet, oldChar) + offset) % len(alphabet)
		newChar := string(alphabet[newIdx])

		encodedMsg += newChar
	}
	
	return encodedMsg
}

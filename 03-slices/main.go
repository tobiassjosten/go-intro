package main

import "fmt"

func main() {
	messages := []string{
		"hello world!",
		"hello vars!",
		"hello slices!",
	}

	for _, message := range messages {
		fmt.Println(message)
	}
	// messages[2:3]
	// for { }

	// []byte{'h', 'e', 'l', 'l', 'o']
}

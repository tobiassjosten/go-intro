package main

import "fmt"

func main() {
	messages := map[string]string{
		"first":  "hello world!",
		"second": "hello vars!",
		"third":  "hello maps!",
	}

	for count, message := range messages {
		fmt.Printf("%s: %s\n", count, message)
	}
	// messages["third"]
	// `go run` x10 – order
}

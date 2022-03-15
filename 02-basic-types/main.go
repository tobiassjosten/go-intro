package main

import "fmt"

func main() {
	var message string = "hello vars"

	punctuation := "!"

	fmt.Println(message + punctuation)

	message = 1337

	fmt.Printf("%s%s\n", message, punctuation)
}

// func message() (message string, punctuation string)

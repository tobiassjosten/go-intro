package main

import "fmt"

type Coder struct {
	Name string
}

func main() {
	coder := Coder{"Tobias"}
	fmt.Printf("hello %s!\n", coder.Name)
}

// .Age
// Coder{"Tobias", 37} / Coder{Name: "Tobias"}

// .Message()

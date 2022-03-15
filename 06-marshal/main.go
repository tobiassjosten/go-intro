package main

import (
	"encoding/json"
	"fmt"
)

// {"name":"Tobias","age":37}

type Coder struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	coder := Coder{"Tobias", 37}

	data, _ := json.Marshal(coder)

	fmt.Println(data)
	// string(data)

	// _ = json.Unmarshal(data, &coder2)
}

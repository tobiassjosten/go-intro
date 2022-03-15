package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Coder struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	coder := Coder{Name: "Tobias"}

	resp, _ := http.Get("https://api.agify.io/?name=" + coder.Name)

	body, _ := io.ReadAll(resp.Body)

	_ = json.Unmarshal(body, &coder)

	fmt.Println(coder.Age)
}

// func age()

// `time go run`

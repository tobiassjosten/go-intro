package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func main() {
	name := "Tobias"

	ch := make(chan string)

	go age(name, ch)
	go gender(name, ch)
	go country(name, ch)
}

func age(name string, ch chan string) {
	resp, _ := http.Get("https://api.agify.io/?name=" + name)

	body, _ := io.ReadAll(resp.Body)

	var data struct {
		Age int `json:"age"`
	}
	_ = json.Unmarshal(body, &data)

	fmt.Println("age", data.Age)
}

func gender(name string, ch chan string) {
	resp, _ := http.Get("https://api.genderize.io/?name=" + name)

	body, _ := io.ReadAll(resp.Body)

	var data struct {
		Gender int `json:"gender"`
	}
	_ = json.Unmarshal(body, &data)

	fmt.Println("gender", data.Gender)
}

func country(name string, ch chan string) {
	resp, _ := http.Get("https://api.nationalize.io/?name=" + name)

	body, _ := io.ReadAll(resp.Body)

	var data struct {
		Countries []struct {
			CountryID string `json:"country_id"`
		} `json:"country"`
	}
	_ = json.Unmarshal(body, &data)

	fmt.Println("country", data.Countries)
}

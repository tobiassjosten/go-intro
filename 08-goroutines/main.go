package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func main() {
	name := "Tobias"
	age(name)
	gender(name)
	country(name)

	// go age(name)
	// time.Sleep(5 * time.Second)

	// var wg sync.WaitGroup
	// .Add(1) .Done() .Wait()
}

func age(name string) {
	resp, _ := http.Get("https://api.agify.io/?name=" + name)

	body, _ := io.ReadAll(resp.Body)

	var data struct {
		Age int `json:"age"`
	}
	_ = json.Unmarshal(body, &data)

	fmt.Println("age", data.Age)
}

func gender(name string) {
	resp, _ := http.Get("https://api.genderize.io/?name=" + name)

	body, _ := io.ReadAll(resp.Body)

	var data struct {
		Gender int `json:"gender"`
	}
	_ = json.Unmarshal(body, &data)

	fmt.Println("gender", data.Gender)
}

func country(name string) {
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

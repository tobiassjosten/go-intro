package main

import (
	"fmt"
	"log"

	_ "github.com/joho/godotenv/autoload"
	"github.com/plaid/go-envvar/envvar"
)

type Config struct {
	MySecretKey string `envvar:"MY_SECRET_KEY"`
	HiddenToken string `envvar:"HIDDEN_TOKEN" default:"invisible"`
}

func NewConfig() (*Config, error) {
	cfg := &Config{}

	if err := envvar.Parse(cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}

func main() {
	cfg, err := NewConfig()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(cfg.MySecretKey, cfg.HiddenToken)
}

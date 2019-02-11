package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Config struct {
	Redirects []Redirect `json:"redirects"`
}

type Redirect struct {
	From string `json:"from"`
	To   string `json:"to"`
}

func ReadConfig() *Config {
	configJSON, err := ioutil.ReadFile("./config.json")
	jsonBytes := ([]byte)(configJSON)
	if err != nil {
		panic(err)
	}
	data := new(Config)
	err = json.Unmarshal(jsonBytes, data)
	if err != nil {
		fmt.Println("error:", err)
	}

	return data
}

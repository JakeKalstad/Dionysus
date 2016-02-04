package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type Configuration struct {
}

func (c Configuration) readConfig() {
	file, _ := os.Open("config.json")
	decoder := json.NewDecoder(file)
	configuration := Configuration{}
	err := decoder.Decode(&configuration)
	if err != nil {
		fmt.Println("error:", err)
	}
}

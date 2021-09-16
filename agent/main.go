package main

import (
	"encoding/json"
	"fmt"
	"time"

	"./app"
)

func main() {
	config, err := LoadEmbeddedConfig()
	// config, err := LoadConfig("config.json")
	if err != nil {
		fmt.Printf("Error reading config file: %s\n", err)
		return
	}

	config_json, _ := json.Marshal(&config)
	fmt.Printf("Loaded config:\n%s\n", config_json)

	client, err := CreateClient(config)
	if err != nil {
		fmt.Printf("Error creating client: %s\n", err)
		return
	}

	agent := app.NewAgent(client)

	for {
		err := agent.Heartbeat()
		if err != nil {
			fmt.Printf("Error: %s\n", err)
		}
		time.Sleep(time.Duration(config.IntervalMs) * time.Millisecond)
	}

}

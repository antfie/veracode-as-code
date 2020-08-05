package main

import "encoding/json"

func parseJsonConfig(data []byte) Config {
	var c Config

	err := json.Unmarshal(data, &c)

	if err != nil {
		panic(err)
	}

	return c
}

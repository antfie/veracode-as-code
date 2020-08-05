package main

import "github.com/go-yaml/yaml"

func parseYamlConfig(data []byte) Config {
	var c Config

	err := yaml.Unmarshal(data, &c)

	if err != nil {
		panic(err)
	}

	return c
}

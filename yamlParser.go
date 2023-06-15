package main

import "gopkg.in/yaml.v3"

func parseYamlConfig(data []byte) Config {
	var c Config

	err := yaml.Unmarshal(data, &c)

	if err != nil {
		panic(err)
	}

	return c
}

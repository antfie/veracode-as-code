package main

import (
	"path/filepath"
	"strings"
)

func loadConfig() Config {
	files, err := filepath.Glob("scan.*")

	if err != nil {
		panic("Could not find a scan file.")
	}

	for _, f := range files {
		if strings.ToLower(f) == "veracode.yaml" {
			configFile, err := readFile(f)

			if err != nil {
				panic(err)
			}

			return parseYamlConfig(configFile)
		}
	}

	panic("Could not find a scan file.")
}

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
		if strings.ToLower(f) == "scan.json" {
			configFile, err := readFile(f)

			if err != nil {
				panic(err)
			}

			return parseYamlConfig(configFile)
		}

		if strings.ToLower(f) == "scan.yml" || strings.ToLower(f) == "scan.yaml" {
			configFile, err := readFile(f)

			if err != nil {
				panic(err)
			}

			return parseYamlConfig(configFile)
		}
	}

	panic("Could not find a scan file.")
}

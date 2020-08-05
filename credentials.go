package main

import (
	"bufio"
	"os"
	"path"
	"strings"
)

func getCredentials() (string, string) {
	id := os.Getenv("VERACODE_API_KEY_ID")
	secret := os.Getenv("VERACODE_API_KEY_SECRET")

	if len(id) > 0 && len(secret) > 0 {
		return id, secret
	}

	home, err := os.UserHomeDir()

	if err != nil {
		panic(err)
	}

	file, err := os.Open(path.Join(home, ".veracode", "credentials"))

	if err != nil {
		panic(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	found := false

	for scanner.Scan() {
		line := scanner.Text()

		if line == "[default]" {
			found = true
		} else if found {
			if strings.Contains(line, "veracode_api_key_id = ") {
				id = strings.Split(line, "veracode_api_key_id = ")[1]
			}

			if strings.Contains(line, "veracode_api_key_secret = ") {
				secret = strings.Split(line, "veracode_api_key_secret = ")[1]
			}
		}

		if len(id) > 0 && len(secret) > 0 {
			return id, secret
		}
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	panic("No credentials found")
}

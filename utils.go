package main

import (
	"io/ioutil"
	"os"
)

func readFile(filename string) ([]byte, error) {
	file, err := os.Open(filename)

	if err != nil {
		return nil, err
	}

	defer file.Close()

	data, err := ioutil.ReadAll(file)

	if err != nil {
		return nil, err
	}

	return data, nil
}

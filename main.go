package main

import (
	"log"

	api "github.com/antfie/veracode-go-api"
)

func main() {
	accessKeyID, secretAccessKey := getCredentials()
	config := loadConfig()

	log.Printf("Application: %s", config.Application.Name)

	applications, err := api.GetApplicationList(accessKeyID, secretAccessKey)

	if err != nil {
		panic(err)
	}

	log.Print(applications.Applications)
}

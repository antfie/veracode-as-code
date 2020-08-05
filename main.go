package main

import (
	"log"

	api "github.com/antfie/veracode-go-api"
)

func main() {
	accessKeyID, secretAccessKey := getCredentials()
	config := loadConfig()
	
	applications, err := api.GetApplicationList(accessKeyID, secretAccessKey)

	if err != nil {
		panic(err)
	}

	log.Print(applications.Applications)
}

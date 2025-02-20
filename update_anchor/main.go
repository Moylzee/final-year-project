package main

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"
	"time"
)

var (
	NewSwaggerFilePath    = "../Bucket/latest_swagger/latest_swagger.json"
	AnchorSwaggerFilePath = "../Bucket/anchor_swagger/anchor.json"
	HistFilePath          = "../Bucket/hist/swagger/"
)

func main() {
	date := time.Now().Format("02-01-2006")

	log.Println("Storing Anchor File in Hist")
	// Read the swagger file
	data, err := os.ReadFile(AnchorSwaggerFilePath)
	if err != nil {
		log.Fatalf(err.Error())
	}

	err = os.WriteFile(filepath.Join(HistFilePath, date+".json"), data, 0644)
	if err != nil {
		log.Fatalf(err.Error())
	}

	log.Println("Stored Anchor File in Hist")

	log.Println("Reading Swagger File From File")
	// Read the swagger file
	swag, err := os.ReadFile(NewSwaggerFilePath)
	if err != nil {
		log.Fatalf(err.Error())
	}
	log.Println("Read Swagger File From File")

	log.Println("Overwriting Anchor File")
	err = os.WriteFile(AnchorSwaggerFilePath, swag, 0644)
	if err != nil {
		log.Fatalf(err.Error())
	}
	log.Println("Anchor File Succesfully Overwritten")
}

func WriteToFile(d interface{}, filePath string) {
	data, err := json.MarshalIndent(d, "", "  ")

	if err = os.WriteFile(filePath, data, 0644); err != nil {
		log.Fatal(err)
	}
}

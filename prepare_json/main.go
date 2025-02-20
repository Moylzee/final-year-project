package main

import (
	"encoding/json"
	"log"
	"os"

	"github.com/jeremywohl/flatten"
)

var (
	NewSwaggerFilePath    = "../Bucket/latest_swagger/latest_swagger.json"
	AnchorSwaggerFilePath = "../Bucket/anchor_swagger/anchor.json"
)

func main() {
	log.Println("Flattening Latest Swagger")

	err := flattenFile(NewSwaggerFilePath)
	if err != nil {
		log.Printf("Error flattening latest swagger: %v", err)
	}

	log.Println("Flattened Latest Swagger")
	log.Println("Flattening Anchor Swagger")

	err = flattenFile(AnchorSwaggerFilePath)
	if err != nil {
		log.Printf("Error flattening anchor swagger: %v", err)
	}

	log.Println("Flattened Anchor Swagger")
}

func flattenFile(filePath string) error {
	log.Println("Reading Swagger File From File")
	// Read the swagger file
	data, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}

	// Parse JSON
	var swagger map[string]interface{}
	err = json.Unmarshal(data, &swagger)
	if err != nil {
		return err
	}
	log.Println("Read Swagger File From File")

	flattenedSchema, err := flatten.Flatten(swagger, "", flatten.DotStyle)
	if err != nil {
		return err
	}

	if err := WriteToFile(flattenedSchema, filePath); err != nil {
		return err
	}

	return nil
}

func WriteToFile(d interface{}, filePath string) error {
	data, err := json.MarshalIndent(d, "", "  ")

	if err = os.WriteFile(filePath, data, 0644); err != nil {
		return err
	}
	return nil
}

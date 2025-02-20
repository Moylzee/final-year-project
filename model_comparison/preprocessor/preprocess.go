package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strings"
)

var (
	resultsFilepath = "../../bucket/comparison_results/results.csv"
	objects         []string
)

func main() {
	err := preprocessData()
	if err != nil {
		fmt.Printf("Error preprocessing data: %v\n", err)
	}
}

func preprocessData() error {
	// Open the file
	file, err := os.Open(resultsFilepath)
	if err != nil {
		return fmt.Errorf("error opening file: %v", err)
	}
	defer file.Close()

	// Create a new CSV reader
	reader := csv.NewReader(file)

	// Read all records
	records, err := reader.ReadAll()
	if err != nil {
		return fmt.Errorf("error reading CSV: %v", err)
	}

	// Process each record
	for i, r := range records {
		if i == 0 { // Skip header
			continue
		}

		a := strings.Split(r[0], ".")
		objectName := a[0]

		if !contains(objects, objectName) {
			objects = append(objects, objectName)
		}

		// Replace ".properties." with "." if present
		records[i][0] = strings.ReplaceAll(records[i][0], ".properties.", ".")

		if r[0] == fmt.Sprintf("%s.type", objectName) && r[3] == "added" && r[1] == "" {
			log.Println("New Object Detected")
			
		}
	}



	log.Println(objects)
	return nil
}

func contains(slice []string, str string) bool {
	for _, s := range slice {
		if s == str {
			return true
		}
	}
	return false
}

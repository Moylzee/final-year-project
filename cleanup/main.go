package main

import (
	"log"
	"os"
	"time"
)

var (
	date = time.Now().Format("02-01-2006")

	summaryHistDir = "../bucket/hist/summary/"
	summaryDir     = "../bucket/summary/"

	latestSwaggerHistDir = "../bucket/hist/swagger/"
	latestSwaggerDir     = "../bucket/latest_swagger/"
)

func main() {
	log.Println("Swagger Comparison Process Succesfully Completed!")
	log.Println("Beginning Cleanup Process")

	log.Println("Deleting Summary Output")
	if err := deleteSummaryOutput(); err != nil {
		log.Println(err)
	}

	log.Println("Deleting Latest Swagger")
	if err := deleteLatestSwagger(); err != nil {
		log.Println(err)
	}

	log.Println("Cleanup Process Complete!")
}

func deleteSummaryOutput() error {
	log.Println("Checking That Summary is saved in Archive")
	_, err := os.Stat(summaryHistDir + date + ".md")
	if err != nil {
		return err
	}

	log.Println("Summary is in Archive - Continuing Cleanup Process")
	if err := os.Remove(summaryDir + "summary_output.md"); err != nil {
		return err
	}

	return nil
}

func deleteLatestSwagger() error {
	log.Println("Checking That Latest Swagger is saved in Archive")
	_, err := os.Stat(latestSwaggerHistDir + date + ".json")
	if err != nil {
		return err
	}

	log.Println("Latest Swagger is in Archive - Continuing Cleanup Process")
	if err := os.Remove(latestSwaggerDir + "latest_swagger.json"); err != nil {
		return err
	}

	return nil
}

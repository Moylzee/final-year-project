package test

import (
	"log"
	"net/http"
	"testing"

	"gotest.tools/v3/assert"
)

// Test to Ensure URL is valid
func TestAccGetSwaggerFromUrl(t *testing.T) {
	SwaggerUrl := "https://s3.dualstack.us-east-1.amazonaws.com/inin-prod-api/us-east-1/public-api-v2/swagger-schema/publicapi-v2-latest.json"

	log.Println("Retrieving Swagger File From URL")
	resp, err := http.Get(SwaggerUrl)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		log.Fatal("Failed to retrieve Swagger File")
	}

	log.Println("Retrieved Swagger File")

	assert.Equal(t, resp.StatusCode, 200)
}

package main

import (
	"encoding/json"
	"log"
	"os"
	"testing"
)

var (
	expectedRefsFilepath = "../bucket/testing/expected_refs.json"
	testSwaggerFilepath  = "../bucket/testing/test_swagger.json"
	objects              = []string{}
)

func TestUnitAccReferenceObjects(t *testing.T) {
	data, err := os.ReadFile(testSwaggerFilepath)
	if err != nil {
		log.Printf("Error reading swagger file: %v", err)
	}

	// Parse JSON
	var swagger map[string]interface{}
	err = json.Unmarshal(data, &swagger)
	if err != nil {
		log.Printf("Error parsing swagger JSON: %v", err)
	}

	AllRefs(swagger)

	// Read expected refs file
	expectedRefs, err := os.ReadFile(expectedRefsFilepath)
	if err != nil {
		log.Printf("Error reading expected refs file: %v", err)
	}

	// Parse JSON
	var expectedRefsJson []string
	err = json.Unmarshal(expectedRefs, &expectedRefsJson)
	if err != nil {
		log.Printf("Error parsing expected refs JSON: %v", err)
	}

	for _, ref := range expectedRefsJson {
		if !Contains(AllObjects, ref) {
			t.Errorf("Expected %s to be in objects", ref)
		}
	}
	for _, ref := range AllObjects {
		if !Contains(expectedRefsJson, ref) {
			t.Errorf("Expected %s to be in expectedRefsJson", ref)
		}
	}
}

package main

import (
	"bytes"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/enescakir/emoji"
)

var (
	resultsFilePath         = "../Bucket/comparison_results/results.json"
	destinationFilePath     = "../Bucket/summary/summary_output.md"
	histDestinationFilePath = "../Bucket/hist/summary/"

	url        = "https://apps.mypurecloud.com:443/webhooks/api/v1/webhook/34eb455d-cf9c-4bdf-a8b4-1b2418cdb15b?chatGroupId=66852b7db2f6b0e43c9309b4"
	TestUrl    = "https://apps.mypurecloud.com:443/webhooks/api/v1/webhook/34eb455d-cf9c-4bdf-a8b4-1b2418cdb15b"
	summaryURL = "https://github.com/Moylzee/FYP/blob/main/FYP/bucket/summary/summary_output.md"

	date        = time.Now().Format("02-01-2006")
	numAdded    int
	numModified int
	numRemoved  int

	noChangesDetectedMessage = fmt.Sprintf("```\n%s%s%s\nNo Changes Detected In The Swagger\n%s%s%s\n```", emoji.Prohibited, emoji.Prohibited, emoji.Prohibited, emoji.Prohibited, emoji.Prohibited, emoji.Prohibited)
)

func main() {
	log.Println("Starting summary process...")

	if !checkForDifferences() {
		log.Println("No Differences Detected in Swagger")
		//postToChat(noChangesDetectedMessage)
	} else {
		var summaries []string

		log.Println("Reading Summary File")
		data, err := os.ReadFile(resultsFilePath)
		if err != nil {
			panic(err)
		}

		// Unmarshal the JSON into a map[string]interface{}
		var jsonData map[string]interface{}
		err = json.Unmarshal(data, &jsonData)
		if err != nil {
			panic(err)
		}

		// Print each attribute individually
		for key, value := range jsonData {
			var a = fmt.Sprintf("# %s", key)
			b := value.(map[string]interface{})
			added := b["added"]
			removed := b["removed"]
			oldDesc := b["old_description"]
			newDesc := b["new_description"]

			if strings.Contains(key, ".description") {
				a = a + fmt.Sprintf("\n## Old Description\n%v\n## New Description\n%v", oldDesc, newDesc)
			} else {
				if added != nil {
					a = a + "\n## Added"
					for _, item := range added.([]interface{}) {
						a = a + fmt.Sprintf("\n- %v", item)
					}
				}
				if removed != nil {
					a = a + "\n## Removed"
					for _, item := range removed.([]interface{}) {
						a = a + fmt.Sprintf("\n- %v", item)
					}
				}
			}
			summaries = append(summaries, a)
		}

		// Create a buffer for pretty-printed JSON
		var prettyJSON bytes.Buffer
		err = json.Indent(&prettyJSON, data, "", "\t")
		if err != nil {
			panic(err)
		}

		log.Println("Writing Summary To File")
		writeSummaryToFile(summaries)
		log.Println("Summary Written To File")
		//postToChat(buildDetectedMessage(numAdded, numRemoved, numModified))
	}
}

func writeSummaryToFile(summaries []string) {
	outputFile, err := os.Create(destinationFilePath)
	if err != nil {
		panic(err)
	}
	defer outputFile.Close()

	for _, summary := range summaries {
		_, err = outputFile.WriteString(fmt.Sprintf("\n%s\n", summary))
		if err != nil {
			panic(err)
		}
	}
}

func buildDetectedMessage(added, removed, modified int) string {
	return fmt.Sprintf("### Changes Detected in Swagger \n"+
		"```\n[Added: %d] [Removed: %d] [Modified: %d]\n```\n"+
		"View the full changes [HERE](%s)", added, removed, modified, summaryURL)
}

func postToChat(messageBody string) {
	log.Println("Posting Message to Chat Room")

	headers := buildHeaders()
	message := map[string]string{"message": messageBody}
	client := &http.Client{}

	// Convert the message map to JSON
	data, err := json.Marshal(message)
	if err != nil {
		fmt.Printf("Error marshaling JSON: %v\n", err)
		return
	}

	// Create a new request
	req, err := http.NewRequest("POST", TestUrl, nil)
	if err != nil {
		fmt.Printf("Error creating request: %v\n", err)
		return
	}

	// Set headers
	for k, v := range headers {
		req.Header.Set(k, v)
	}

	// Set request body
	req.Body = io.NopCloser(bytes.NewReader(data))

	// Execute the request
	log.Println("Executing Request")
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("Error sending request: %v\n", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode == 200 {
		fmt.Println("Message sent successfully!")
	} else {
		fmt.Printf("Error sending message. Status code: %d\n", resp.StatusCode)
		_, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Printf("Error reading response body: %v\n", err)
			return
		}
	}
}

func buildHeaders() map[string]string {
	headers := map[string]string{
		"Content-Type": "application/json",
		"acceptType":   "application/json"}

	return headers
}

func checkForDifferences() bool {
	file, err := os.Open(resultsFilePath)
	if err != nil {
		log.Fatalf("Error opening file: %v", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)

	// Read the first line
	_, err = reader.Read()
	if err == io.EOF {
		return false // File is empty
	}
	if err != nil {
		log.Fatalf("Error reading record: %v", err)
	}

	// Return true if the file contains stuff
	return true
}

package tools

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

type Result struct {
	RequestIndex int         `json:"request_index"`
	RequestBody  interface{} `json:"request_body"` // Use interface{} to handle any JSON value
	Status       string      `json:"status"`
	ResponseBody interface{} `json:"response_body"` // Use interface{} to handle any JSON value
}

func TestGivenApi(api string, body []map[string]interface{}, interval int) error {
	var results []Result

	for i, v := range body {
		bodyBytes, err := json.Marshal(v)
		if err != nil {
			return err
		}
		fmt.Println("Testing body: ", string(bodyBytes))
		// Create and send the request
		request, err := http.NewRequest("POST", api, bytes.NewBuffer(bodyBytes))
		if err != nil {
			return err
		}
		request.Header.Set("Content-Type", "application/json")

		client := &http.Client{}
		response, err := client.Do(request)
		if err != nil {
			return err
		}
		defer response.Body.Close()

		// Read and process the response
		respBody, err := io.ReadAll(response.Body)
		if err != nil {
			return err
		}

		// Unmarshal request and response bodies
		var reqBody, respBodyFormatted interface{}
		if err := json.Unmarshal(bodyBytes, &reqBody); err != nil {
			return err
		}
		if err := json.Unmarshal(respBody, &respBodyFormatted); err != nil {
			return err
		}

		// Collect the result
		result := Result{
			RequestIndex: i,
			RequestBody:  reqBody,
			Status:       response.Status,
			ResponseBody: respBodyFormatted,
		}
		results = append(results, result)

		// Print response details
		fmt.Printf("Status: %s\n", response.Status)
		fmt.Printf("Response body: %s\n", string(respBody))

		// Sleep for the specified interval
		if interval > 0 {
			fmt.Printf("Waiting for %d seconds...\n", interval)
			time.Sleep(time.Duration(interval) * time.Second)
		}
	}

	// Write results to result.json with pretty formatting
	file, err := os.Create("result.json")
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	err = encoder.Encode(results)
	if err != nil {
		return err
	}

	return nil
}

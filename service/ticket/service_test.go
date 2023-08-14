package ticket

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"testing"
)

func Test_AddJiraTicket(t *testing.T) {
	baseURL := "https://dungcda.atlassian.net/rest/api/2"
	apiToken := "Y2FvZHVuZzU2NkBnbWFpbC5jb206QVRBVFQzeEZmR0YwYkVzazhjWVJSOE9nTEs0VXJzRUU3SW9VOWVINDBFaVhCTkFEdHRJalhGb1ZnZXpWSm9vLVhDdG5HcWt4d1pwdlh5TkdlZS1fVXRxVjZVRjc5YjBodTdPbmt0b3RCVERTWVNZZy1hREc5bVNoQWZBTGRnVjVYZERRaHV3ZFZfWXQtb2RpRFVtQTFNZE1KMG9FdW1tejQ0bHVRTWVRcnliMEZ3aEtFN2x3SFRFPThEMDNGRDE3"

	// Create a JSON payload for the Jira ticket
	payload := map[string]interface{}{
		"fields": map[string]interface{}{
			"project": map[string]interface{}{
				"key": "TEK",
			},
			"summary":     "Example Jira Ticket",
			"description": "This is an example Jira ticket created via API.",
			"issuetype": map[string]interface{}{
				"name": "Task",
			},
		},
	}

	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return
	}

	// Create a new HTTP request
	req, err := http.NewRequest("POST", baseURL+"/issue", bytes.NewBuffer(payloadBytes))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Basic "+apiToken)

	// Perform the request
	client := http.DefaultClient
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	fmt.Println(string(body))

	// Check the response
	if resp.StatusCode == http.StatusCreated {
		fmt.Println("Jira ticket created successfully!")
	} else {
		fmt.Println("Failed to create Jira ticket. Status code:", resp.StatusCode)
	}
}

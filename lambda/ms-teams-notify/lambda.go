package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func lambdaHandler(event interface{}, context interface{}) {
	// List of webhook URLs - replace with actual webhooks from MS Teams
	webhookURLs := []string{
		"https://your-webhook.webhook.office.com",
		"https://your-webhook.webhook.office.com",
	}

	httpClient := &http.Client{}

	for _, url := range webhookURLs {
		msg := map[string]interface{}{
			"text": event.(map[string]interface{})["Records"].([]interface{})[0].(map[string]interface{})["Sns"].(map[string]interface{})["Message"],
		}

		encodedMsg, _ := json.Marshal(msg)
		resp, err := httpClient.Post(url, "application/json", bytes.NewBuffer(encodedMsg))
		if err != nil {
			fmt.Println(err)
			return
		}
		defer resp.Body.Close()

		fmt.Println(map[string]interface{}{
			"message":     event.(map[string]interface{})["Records"].([]interface{})[0].(map[string]interface{})["Sns"].(map[string]interface{})["Message"],
			"status_code": resp.Status,
			"response":    resp.Body,
		})
	}
}

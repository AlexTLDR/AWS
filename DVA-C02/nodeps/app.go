package main

import (
	"log"
)

func main() {
	logger := log.New(log.Writer(), "", log.LstdFlags|log.Lmicroseconds|log.LUTC)
	logger.SetOutput(log.Writer())
	logger.SetFlags(log.LstdFlags | log.Lmicroseconds | log.LUTC)

	myHandler(nil, nil)
}

func myHandler(event map[string]interface{}, context interface{}) map[string]interface{} {
	logger := log.Default()
	logger.Printf("Event: %v\n", event)
	logger.Printf("Context: %v\n", context)

	response := map[string]interface{}{
		"statusCode": 200,
		"body": map[string]interface{}{
			"message": "Hello World!",
		},
	}

	return response
}

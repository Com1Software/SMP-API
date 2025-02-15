package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"
)

func main() {
	pdsHost := os.Getenv("PDSHOST")
	accessJWT := os.Getenv("ACCESS_JWT")
	blueskyHandle := os.Getenv("BLUESKY_HANDLE")

	record := map[string]string{
		"text":      "Hello world! I posted this via the API.",
		"createdAt": time.Now().UTC().Format(time.RFC3339),
	}

	data := map[string]interface{}{
		"repo":       blueskyHandle,
		"collection": "app.bsky.feed.post",
		"record":     record,
	}

	payload, err := json.Marshal(data)
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		return
	}

	req, err := http.NewRequest("POST", pdsHost+"/xrpc/com.atproto.repo.createRecord", bytes.NewBuffer(payload))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	req.Header.Set("Authorization", "Bearer "+accessJWT)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("Error: received status code", resp.StatusCode)
		return
	}

	fmt.Println("Success! Record created.")
}

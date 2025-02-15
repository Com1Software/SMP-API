package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func main() {
	url := "https://bsky.app/xrpc/com.atproto.repo.createRecord"
	accessJWT := "YOUR_ACCESS_JWT"
	blueskyHandle := "YOUR_BLUESKY_HANDLE"

	record := map[string]interface{}{
		"text":      "Hello world! I posted this via the API.",
		"createdAt": time.Now().UTC().Format(time.RFC3339),
	}
	data := map[string]interface{}{
		"repo":       blueskyHandle,
		"collection": "app.bsky.feed.post",
		"record":     record,
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		return
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}
	req.Header.Set("Authorization", "Bearer "+accessJWT)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making request:", err)
		return
	}
	defer resp.Body.Close()

	fmt.Println("Response Status:", resp.Status)
}

package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/go-resty/resty/v2"
)

func main() {
	client := resty.New()

	// Replace these with your actual credentials and endpoints
	loginEndpoint := "https://bsky.social/xrpc/com.atproto.server.createSession"
	postEndpoint := "https://bsky.social/xrpc/com.atproto.feed.createPost"
	username := "yourname" // Add your username
	password := "password" // Add your password
	message := "Hello, Bluesky!"

	// Step 1: Obtain a session token
	loginResponse, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(map[string]string{
			"identifier": username,
			"password":   password,
		}).
		Post(loginEndpoint)

	if err != nil {
		log.Fatalf("Error obtaining token: %v", err)
	}

	if loginResponse.StatusCode() != 200 {
		log.Fatalf("Failed to obtain token: %s", loginResponse.String())
	}

	// Parse the session token from the response
	var result map[string]interface{}
	if err := json.Unmarshal(loginResponse.Body(), &result); err != nil {
		log.Fatalf("Error parsing token: %v", err)
	}

	sessionToken, ok := result["token"].(string)
	if !ok {
		log.Fatalf("Invalid token format")
	}

	// Step 2: Use the session token to post a message
	postResponse, err := client.R().
		SetAuthToken(sessionToken). // Use the correct method to set the token
		SetHeader("Content-Type", "application/json").
		SetBody(map[string]string{
			"text": message, // Assuming the API expects a field named "text"
		}).
		Post(postEndpoint)

	if err != nil {
		log.Fatalf("Error posting message: %v", err)
	}

	fmt.Println("Response Status Code:", postResponse.StatusCode())
	fmt.Println("Response Body:", string(postResponse.Body()))
}

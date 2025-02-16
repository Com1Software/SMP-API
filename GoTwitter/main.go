package main

import (
	"fmt"
	"log"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
)

func main() {
	// Replace these with your actual API credentials
	consumerKey := "your_consumer_key"
	consumerSecret := "your_consumer_secret"
	accessToken := "your_access_token"
	accessSecret := "your_access_secret"

	config := oauth1.NewConfig(consumerKey, consumerSecret)
	token := oauth1.NewToken(accessToken, accessSecret)
	httpClient := config.Client(oauth1.NoContext, token)

	client := twitter.NewClient(httpClient)

	// Home Timeline
	tweets, _, err := client.Timelines.HomeTimeline(&twitter.HomeTimelineParams{
		Count: 10,
	})
	if err != nil {
		log.Fatalf("Error getting timeline: %v", err)
	}

	for _, tweet := range tweets {
		fmt.Printf("User: %s\nTweet: %s\n\n", tweet.User.ScreenName, tweet.Text)
	}
}

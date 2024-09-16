package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/slack-go/slack"
)

func main() {
	// Load environment variables from the .env file
	err := godotenv.Load()
	if err != nil {
		fmt.Printf("Error loading .env file: %v\n", err)
		return
	}

	// Get Slack Bot Token and Channel ID from environment variables
	slackBotToken := os.Getenv("SLACK_BOT_TOKEN")
	slackChannelID := os.Getenv("SLACK_CHANNEL_ID")

	if slackBotToken == "" || slackChannelID == "" {
		fmt.Println("Missing required environment variables")
		return
	}

	// Initialize Slack API client
	api := slack.New(slackBotToken)

	// Define the file to upload
	filePath := "quote.png" // Replace with your file path

	// Upload parameters
	fileUploadParams := slack.FileUploadParameters{
		File:     filePath,
		Channels: []string{slackChannelID},
	}

	// Upload file
	_, err = api.UploadFile(fileUploadParams)
	if err != nil {
		fmt.Printf("Error uploading file: %s\n", err.Error())
	} else {
		fmt.Println("File uploaded successfully!")
	}
}

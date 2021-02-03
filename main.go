package main

import (
	"log"
	"os"

	"github.com/nikoksr/notify"
	"github.com/nikoksr/notify/service/slack"
)

func main() {

	// get token env
	token := os.Getenv("SLACK_OAUTH_TOKEN")
	if len(token) <= 0 {
		log.Println("Enviroment SLACK_OAUTH_TOKEN not fount")
		return
	}

	// get channel env
	channel := os.Getenv("SLACK_CHANNEL_ID")
	if len(token) <= 0 {
		log.Println("Enviroment SLACK_CHANNEL_ID not fount")
		return
	}

	// Start a new instance
	notifier := notify.New()

	// Provide your Slack OAuth Access Token
	slackService := slack.New(token)

	// Passing a Slack channel id as receiver for our messages.
	// Where to send our messages.
	// https://xxxx.slack.com/archives/C0NE3CG4D
	slackService.AddReceivers(channel)

	// Tell our notifier to use the Slack service. You can repeat the above process
	// for as many services as you like and just tell the notifier to use them.
	notifier.UseService(slackService)

	// Send a message
	if err := notifier.Send(
		"Hello :wave:\n",
		"Here your stats for today!\nAccess report -> 1.1.1.1\nDNS server: xyz.com",
	); err != nil {
		log.Fatal(err)
	}
}

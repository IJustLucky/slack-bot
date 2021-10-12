package main

import (
	"log"

	"github.com/slack-go/slack"
)

func main() {
	OAUTH_TOKEN := "xoxb-2490105638976-2605767109921-TFhELbrhm6kBFeT1Kco7mnSZ"
	CHANNEL_ID := "C02DWLB0P44"

	api := slack.New(OAUTH_TOKEN)
	attachment := slack.Attachment{
		Pretext: "Pretext",
		Text:    "Hello from GolangDocs!",
	}

	channelId, timestamp, err := api.PostMessage(
		CHANNEL_ID,
		slack.MsgOptionText("This is the main message", false),
		slack.MsgOptionAttachments(attachment),
		slack.MsgOptionAsUser(true),
	)

	if err != nil {
		log.Fatalf("%s\n", err)
	}
	log.Printf("Mesagge sucessfully sent to Channel %s\n", channelId, timestamp)
}

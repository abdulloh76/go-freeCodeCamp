package main

import (
	"fmt"
	"os"

	"github.com/slack-go/slack"
)

func main() {
	os.Setenv("SLACK_BOT_TOKEN", "xoxb-2422907983831-3492811706341-jTiUihroGhuSlXZdqa5RcemL")
	os.Setenv("CHANNEL_ID", "C02D87DBD4Z")

	api := slack.New(os.Getenv("SLACK_BOT_TOKEN"))
	channelArr := []string{os.Getenv("CHANNEL_ID")}
	fileArr := []string{"linkedincv.pdf"}

	for _, file := range fileArr {
		params := slack.FileUploadParameters{
			Channels: channelArr,
			File:     file,
		}
		uploadedFile, err := api.UploadFile(params)
		if err != nil {
			fmt.Printf("%s\n", err)
			return
		}
		fmt.Printf("Name: %s, URL:%s\n", uploadedFile.Name, uploadedFile.URL)
	}
}

package main

import (
	"os"

	"github.com/shomali11/slacker"
)

func printCommandEvents(analyticsChannel <-chan *slacker.CommandEvent) {}

func main() {
	os.Setenv("SLACK_BOT_TOKEN", "xoxb-2422907983831-3518453594352-EdgenDMQ6oYb51RIMm56yPBD")
	os.Setenv("SLACK_APP_TOKEN", "xapp-1-A03EX91K39P-3494671667411-b661fe4b68b28c0c1394d0ff4c34e1f2ba40cacaaf5a2d2754bde0a7fa1e80df")

	bot := slacker.NewClient(os.Getenv("SLACK_BOT_TOKEN"), os.Getenv("SLACK_APP_TOKEN"))

	go printCommandEvents(bot.CommandEvents())
}

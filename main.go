package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/shomali11/slacker"
)

func printCommandEvents(analyticsChannel <-chan *slacker.CommandEvent) {
	for event := range analyticsChannel {
		fmt.Println("Command Events")
		fmt.Println(event.Timestamp)
		fmt.Println(event.Command)
		fmt.Println(event.Parameters)
		fmt.Println(event.Event)
		fmt.Println()
	}
}

func main() {
	os.Setenv("SLACK_BOT_TOKEN", "xoxb-2422907983831-3518453594352-EdgenDMQ6oYb51RIMm56yPBD")
	os.Setenv("SLACK_APP_TOKEN", "xapp-1-A03EX91K39P-3494671667411-b661fe4b68b28c0c1394d0ff4c34e1f2ba40cacaaf5a2d2754bde0a7fa1e80df")

	bot := slacker.NewClient(os.Getenv("SLACK_BOT_TOKEN"), os.Getenv("SLACK_APP_TOKEN"))

	go printCommandEvents(bot.CommandEvents())

	bot.Command("my year of born is <year>", &slacker.CommandDefinition{
		Description: "year of born calculator",
		Example:     "my year of born is 2000",
		Handler: func(botCtx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter) {
			year := request.Param("year")
			yob, err := strconv.Atoi(year)
			if err != nil {
				println("error: year is invalid")
			}
			actualYear, _, _ := time.Now().Date()
			age := actualYear - yob
			res := fmt.Sprintf("age is %d", age)
			response.Reply(res)
		},
	})

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err := bot.Listen(ctx)

	if err != nil {
		log.Fatal(err)
	}

}

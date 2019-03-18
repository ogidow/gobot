package main

import (
	"os"
	"net/http"

	"github.com/ogidow/gobot"
	"github.com/ogidow/gobot/examples/food"
)

func main() {
	bot := gobot.NewGobot()
	bot.AddMachine(food.NewMachine())

	slackVerificationToken := "YOURE_SLACK_VERIFICATION_TOKEN"
	SlackAccessToken := "YOURE_SLACK_ACCESS_TOKEN"

	eventHandler := bot.NewEventApiHandler(slackVerificationToken, "what do you do?", SlackAccessToken)
	http.Handle("/path/to/your_event_api", eventHandler)

	interactiveHandler := bot.NewInteractiveApiHandler()
	http.Handle("/path/to/your_interactive_messages_api", interactiveHandler)
	http.ListenAndServe(":" + os.Getenv("PORT"), nil)
}

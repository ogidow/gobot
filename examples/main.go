package main

import (
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
	http.Handle("/events", eventHandler)

	interactiveHandler := bot.NewInteractiveApiHandler()
	http.Handle("/interactive_messages", interactiveHandler)
}

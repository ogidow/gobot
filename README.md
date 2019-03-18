# gobot
gobot is a bot framework that uses slack interactive message

## Usage

### Create gobot and Machine
In gobot, a series of interactions are called `Machine`.
gobot can register multiple` Machines`.
gobot can be created in the following way.

```go
bot := gobot.NewGobot()
```

`Machine` can be created in the following way.

```golang
m := machine.NewMachine("machineName")
```

`Machine` consists of multiple` State`.
`State` is a slack message itself, and transitions to the next` State` when an event is fired for `State`.

You can add `State` to` Machine` as follows

```go
m.AddState("asking", func(s *machine.State) {
    s.InitialState()
    s.BuildAttachment(func(ev slack.InteractionCallback) {
        s.Text("Do you drink beer?")
        s.Button("accept", "yes", "yes")
        s.Button("decline", "no", "no")
    })
    s.Event("accept", "ordering", func(ev slack.InteractionCallback){})
    s.Event("decline", "canceling", func(ev slack.InteractionCallback){})
  })
```

Call `s.InitialState ()` for the first `State` of` Machine` and `s.EndState ()` for the last `State`.


You can define a Slack Attachment using DSL in `s.BuildAttachment` method.
You can also define an Event for `State` in` s.Event` method.
In the above example, when the `yes` button is pressed, the` accept` event fires and transitions to the `ordering` State.


### register Machine to gobot
You can register the created machine in gobot by the following method

```go
bot.AddMachine(food.NewMachine())
```

### handle slack api
gobot provides a slakc api handler.
gobot uses the event api 'app_mention' to handle mentions.

You can register a handler in the following way to listen to slack api.

```go
slackVerificationToken := "YOURE_SLACK_VERIFICATION_TOKEN"
SlackAccessToken := "YOURE_SLACK_ACCESS_TOKEN"
eventHandler := bot.NewEventApiHandler(slackVerificationToken, "what do you do?", SlackAccessToken)
http.Handle("/path/to/your_event_api", eventHandler)

interactiveHandler := bot.NewInteractiveApiHandler()
http.Handle("/path/to/your_interactive_messages_api", interactiveHandler)

http.ListenAndServe(":" + os.Getenv("PORT"), nil)
```

### Example
See https://github.com/ogidow/gobot/blob/master/examples

![](./examples/gobot_sample.gif)

# gobot
gobot is a bot framework that uses slack interactive message

## Usage
gobot needs a slack event api's `app_mention`

### create gobot machine
```go
import(
  "github.com/ogidow/gobot"
  "github.com/ogidow/gobot/machine"
)

m := machine.NewMachine("order beer")
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

m.AddState("ordering", func(s *machine.State) {
  s.EndState()
  s.BuildAttachment(func(ev slack.InteractionCallback) {
      s.Text("I received an order beer")
  })
})

m.AddState("canceling", func(s *machine.State) {
  s.EndState()
  s.BuildAttachment(func(ev slack.InteractionCallback) {
      s.Text("I canceled your order")
  })
})
```

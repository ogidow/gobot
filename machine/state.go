package machine

import (
	"github.com/nlopes/slack"
)

type State struct {
	name string
	attachment *slack.Attachment
	events map[string]Event
	initial bool
	End bool
}

type Event struct {
	Name string
	Process func(slack.InteractionCallback)
	To string
}


func NewState(name string) *State {
	events := map[string]Event{}
	return &State{name, &slack.Attachment{Color: "#f9a41b", CallbackID: name}, events, false, false}
}

func (d *State) InitialState() {
	d.initial = true
}

func (d *State) EndState() {
	d.End = true
}

func (d *State) Color(c string) {
	d.attachment.Color = c
}

func (d *State) Text(t string) {
	d.attachment.Text = t
}

func (d *State) Button(name string, text string, value string) {
	button := slack.AttachmentAction {
		Name:  name,
      	Text:  text,
      	Type:  "button",
      	Value: value,
    }

	d.attachment.Actions = append(d.attachment.Actions , button)
}

func (d *State) Field(title string, value string) {
	field := slack.AttachmentField {
		Title: title,
		Value: value,
	}

	d.attachment.Fields = append(d.attachment.Fields, field)
}

func (d *State) Event (name string, to string, e func(slack.InteractionCallback)) {
	d.events[name] = Event{
		Name: name,
		To: to,
		Process: e,
	}
}

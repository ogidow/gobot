package machine

import (
	"github.com/nlopes/slack"
)

type State struct {
	name string
	attachment *slack.Attachment
	Events map[string]Event
	BuildAttachmentFunc func(slack.InteractionCallback)
	initial bool
	End bool
}

type Event struct {
	Name string
	Process func(slack.InteractionCallback)
	To string
}

type SelectBoxOption struct {
	Text string
	Value string
}

func NewState(name string) *State {
	events := map[string]Event{}
	return &State{name, &slack.Attachment{Color: "#f9a41b", CallbackID: name}, events, nil, false, false}
}

func (s *State) InitialState() {
	s.initial = true
}

func (s *State) EndState() {
	s.End = true
}

func (s *State) Color(c string) {
	s.attachment.Color = c
}

func (s *State) Text(t string) {
	s.attachment.Text = t
}

func (s *State) Button(name string, text string, value string) {
	button := slack.AttachmentAction {
		Name:  name,
      	Text:  text,
      	Type:  "button",
      	Value: value,
    }

	s.attachment.Actions = append(s.attachment.Actions , button)
}

func (s *State) Field(title string, value string) {
	field := slack.AttachmentField {
		Title: title,
		Value: value,
	}

	s.attachment.Fields = append(s.attachment.Fields, field)
}

func (s *State) SelectBox(name string, options []SelectBoxOption) {
	var slackOptions []slack.AttachmentActionOption

	for _, op := range options {
		slackOptions = append(slackOptions, slack.AttachmentActionOption{Text: op.Text, Value: op.Value})
	}
	selectBox := slack.AttachmentAction {
		Name:    name,
		Type:    "select",
		Options: slackOptions,
	}

	s.attachment.Actions = append(s.attachment.Actions , selectBox)
}

func (s *State) Event (name string, to string, e func(slack.InteractionCallback)) {
	s.Events[name] = Event{
		Name: name,
		To: to,
		Process: e,
	}
}

func (s *State) BuildAttachment(e func(slack.InteractionCallback)) {
	s.BuildAttachmentFunc = e
}

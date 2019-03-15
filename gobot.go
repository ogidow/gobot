package gobot

import (
	"encoding/json"
	"net/http"

	"github.com/nlopes/slack"
	"github.com/looplab/fsm"
)

type Gobot struct {
	machines map[string]SlackEventMachine
	states map[string]SlackEventMachine
}

type SlackEventMachine interface {
	GetStateMachine() *fsm.FSM
	GetNextSlackAttachments() []slack.Attachment
}

func NewGobot() *Gobot {
	machines := map[string]SlackEventMachine{}
	states := map[string]SlackEventMachine{}
	return &Gobot{machines, states}
}

func (g *Gobot) AddMachine(name string, machine SlackEventMachine) {
	g.machines[name] = machine
}

func (g *Gobot)HandleAndResponse(w http.ResponseWriter, callbackEvent slack.InteractionCallback) {
	action := callbackEvent.Actions[0].Name
	messageTs := callbackEvent.MessageTs
	machine := g.states[messageTs]
	if machine == nil {
		machineName := callbackEvent.Actions[0].SelectedOptions[0].Value
		machine = g.machines[machineName]
		g.states[messageTs] = machine
	}

	machine.GetStateMachine().Event(action, callbackEvent)

	message := slack.Msg{
		ReplaceOriginal: true,
		Attachments:     machine.GetNextSlackAttachments(),
	}

	if machine.GetStateMachine().Current() == "end" {
		delete(g.states, messageTs)
	}

	w.Header().Add("Content-type", "application/json")
	json.NewEncoder(w).Encode(&message)
}

func (g *Gobot)GetMachines() []slack.AttachmentActionOption {
	var machines []slack.AttachmentActionOption

	for name := range g.machines {
		machines = append(machines, slack.AttachmentActionOption{Text: name, Value: name})
	}

	return machines
}

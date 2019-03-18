package gobot

import (
	"encoding/json"
	"net/http"

	"github.com/nlopes/slack"
	"github.com/looplab/fsm"
	"github.com/ogidow/gobot/machine"
)

type Gobot struct {
	machines map[string]machine.Machine
	states map[string]*machine.Machine
}

type SlackEventMachine interface {
	GetStateMachine() *fsm.FSM
	GetNextSlackAttachments() []slack.Attachment
}

func NewGobot() *Gobot {
	machines := map[string]machine.Machine{}
	states := map[string]*machine.Machine{}
	return &Gobot{machines, states}
}

func (g *Gobot) AddMachine(machine machine.Machine) {
	g.machines[machine.Name] = machine
}

func (g *Gobot)HandleAndResponse(w http.ResponseWriter, callbackEvent slack.InteractionCallback) {
	action := callbackEvent.Actions[0].Name
	messageTs := callbackEvent.MessageTs
	machine := g.states[messageTs]
	if machine == nil {
		machineName := callbackEvent.Actions[0].SelectedOptions[0].Value
		tmpMachine := g.machines[machineName]
		g.states[messageTs] = &tmpMachine
		machine = &tmpMachine
	} else {
		machine.Event(action, callbackEvent)
	}

	message := slack.Msg{
		ReplaceOriginal: true,
		Attachments:     []slack.Attachment{machine.Attachment()},
	}

	if machine.Current.End {
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

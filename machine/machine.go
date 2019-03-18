package machine

import (
	"github.com/nlopes/slack"
)

type Machine struct {
	Name string
	states map[string]*State
	Current *State
}

func NewMachine(name string) *Machine{
	states := map[string]*State{}
	return &Machine{Name: name, states: states}
}

func (m *Machine) AddState(stateName string, f func(s *State)) {
	state := NewState(stateName)
	f(state)
	m.states[stateName] = state

	if state.initial {
		m.Current = state
	}
}

func(m *Machine) Event(name string, callback slack.InteractionCallback) {
	ev := m.Current.Events[name]
	ev.Process(callback)
	m.Current = m.states[ev.To]
}

func (m *Machine) Attachment() slack.Attachment{
	return *m.Current.attachment
}

func(m *Machine) BuildAttachment(callback slack.InteractionCallback){
	m.Current.clearAttachment()
	m.Current.BuildAttachmentFunc(callback)
}

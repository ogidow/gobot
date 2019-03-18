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
	return &Machine{Name: name}
}

func (m *Machine) AddState(f func(s *State)) {
	state := NewState()
	f(state)
	m.states[state.name] = state

	if state.initial {
		m.Current = state
	}
}

func(m *Machine) Event(name string, callback slack.InteractionCallback) {
	ev := m.Current.events[name]
	ev.Process(callback)
	m.Current = m.states[ev.To]
}

func (m *Machine) Attachment() slack.Attachment{
	return *m.Current.attachment
}

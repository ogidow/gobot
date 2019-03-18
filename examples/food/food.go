package food

import (
	"fmt"

	"github.com/ogidow/gobot/machine"
	"github.com/nlopes/slack"
)

type Food struct {
	Name string
}

func NewMachine() machine.Machine {
	var food Food
	machine := machine.NewMachine("orderFood")
	machine.AddState("selecting", food.selectingState())
	machine.AddState("cheking", food.chekingState())
	machine.AddState("canceling", food.cancelingState())
	machine.AddState("finish", food.finishState())

	return *machine
}

func (f *Food) selectingState() func(s *machine.State) {
	return func(s *machine.State) {
		s.InitialState()
		s.BuildAttachment(func(ev slack.InteractionCallback) {
			options := []machine.SelectBoxOption {
				{Text: "Pizza", Value: "Pizza"},
				{Text: "Sandwich", Value: "Sandwich"},
				{Text: "Hamburger", Value: "Hamburger"},
			}
			s.Text("with food do you want?")
			s.SelectBox("selectFood", options)
			s.Button("canceling", "cancel", "false")
		})

		s.Event("selectFood", "cheking", func(ev slack.InteractionCallback){
			// writing selectFood event logic
			f.Name = ev.Actions[0].SelectedOptions[0].Value
		})
    	s.Event("canceling", "canceling", func(ev slack.InteractionCallback){
			// writing canceling event logic
		})
	}
}

func (f *Food) cancelingState() func(s *machine.State) {
	return func(s *machine.State) {
		s.EndState()

		s.BuildAttachment(func(ev slack.InteractionCallback) {
			s.Text("I will wait for the next use")
		})
	}
}

func (f *Food) chekingState() func(s *machine.State) {
	return func(s *machine.State) {
		s.BuildAttachment(func(ev slack.InteractionCallback) {
			s.Text(fmt.Sprintf("Okay, so that’s one %s?", f.Name))
			s.Button("accept", "yes", "yes")
			s.Button("decline", "no", "no")
		})

		s.Event("accept", "finish", func(ev slack.InteractionCallback){
			// writing accept logic
		})
    	s.Event("decline", "selecting", func(ev slack.InteractionCallback){
			// writing canceling logic
		})
	}
}

func (f *Food) finishState() func(s *machine.State) {
	return func(s *machine.State) {
		s.EndState()

		s.BuildAttachment(func(ev slack.InteractionCallback) {
			s.Text(fmt.Sprintf("I received an order with %s", f.Name))
		})
	}
}

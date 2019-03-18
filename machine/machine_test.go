package machine

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddState(t *testing.T) {
	target := NewMachine("test")

	target.AddState("fuga", func(s *State) {
		s.Text("hoge")
		s.InitialState()
	})

	assert.NotNil(t, target.states["fuga"])
	assert.NotNil(t, target.Current)
}

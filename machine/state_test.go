package machine

import (
	"testing"

	"github.com/nlopes/slack"
	"github.com/stretchr/testify/assert"
)

func TestInitialState(t *testing.T) {
	target := NewState("test")
	target.InitialState()
	assert.Equal(t, true, target.initial)
}

func TestEndState(t *testing.T) {
	target := NewState("test")
	target.EndState()
	assert.Equal(t, true, target.End)
}

func TestColor(t *testing.T) {
	target := NewState("test")
	target.Color("#ffffff")
	assert.Equal(t, "#ffffff", target.attachment.Color)
}

func TestText(t *testing.T) {
	target := NewState("test")
	target.Text("Title Text")
	assert.Equal(t, "Title Text", target.attachment.Text)
}

func TestButton(t *testing.T) {
	target := NewState("test")
	target.Button("ok", "accept", "1")

	assert.Equal(t, "ok", target.attachment.Actions[0].Name)
	assert.Equal(t, "accept", target.attachment.Actions[0].Text)
	assert.Equal(t, "button", target.attachment.Actions[0].Type)
	assert.Equal(t, "1", target.attachment.Actions[0].Value)
}

func TestField(t *testing.T) {
	target := NewState("test")
	target.Field("field title", "field value")

	assert.Equal(t, "field value", target.attachment.Fields[0].Value)
	assert.Equal(t, "field title",  target.attachment.Fields[0].Title)
}

func TestEvent(t *testing.T) {
	ev := func(slack.InteractionCallback) {

	}

	target := NewState("test")
	target.Event("hoge event", "fuga event", ev)

	assert.Equal(t, "hoge event", target.events["hoge event"].Name)
	assert.Equal(t, "fuga event", target.events["hoge event"].To)
}

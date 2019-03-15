package gobot

import (
	"net/http"
	"encoding/json"

	"github.com/nlopes/slack"
)

type interactiveApiHandler struct {
	bot *Gobot
}

func NewInteractiveApiHandler(bot *Gobot) interactiveApiHandler {
	return interactiveApiHandler{bot}
}

func (h interactiveApiHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	r.ParseForm()
	payload := r.PostForm.Get("payload")
	var callback slack.InteractionCallback
	json.Unmarshal([]byte(payload), &callback)

	h.bot.HandleAndResponse(w, callback)
}

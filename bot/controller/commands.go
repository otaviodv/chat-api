package controller

import (
	"chat-api-bot/commands"
	"chat-api-bot/helper"
	"chat-api-bot/model"
	"encoding/json"
	"errors"
	"net/http"
	"strings"
)

type baseResponse struct {
	Message string `json:"message"`
}

func HandleCommand(w http.ResponseWriter, r *http.Request) {

	cmd := model.CommandTransport{}
	err := json.NewDecoder(r.Body).Decode(&cmd)
	if err != nil {
		helper.ErrorJSON(w, err)
		return
	}

	if len(cmd.Text) < 2 || !strings.HasPrefix(cmd.Text, "/") {
		helper.ErrorJSON(w, errors.New("invalid command"))
		return
	}

	go commands.CheckAndRunCommand(cmd)

	m, _ := json.Marshal(baseResponse{"message sent"})
	w.Write(m)
}

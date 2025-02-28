package controller

import (
	"chat-api/helper"
	"chat-api/model"
	"chat-api/service"
	"net/http"
	"strings"

	"github.com/go-chi/chi/v5"
)

func PostMessage(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	msg := model.Message{}
	err := helper.ReadJSON(r, &msg)

	if err != nil {
		helper.ErrorJSON(w, err, http.StatusUnprocessableEntity)
		return
	}
	room := model.Room{}
	err = room.Get(id)
	if err != nil {
		helper.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}
	msg.RoomId = room.Id

	if strings.HasPrefix(msg.Text, "/") {
		if err = service.SubmitCommand(msg); err != nil {
			helper.ErrorJSON(w, err, http.StatusInternalServerError)
			return
		}
		helper.WriteJSON(w, http.StatusOK, msgResponse{"Command Submitted"})
		return
	}

	_, err = msg.Create()
	if err != nil {
		helper.ErrorJSON(w, err, http.StatusInternalServerError)
		return
	}

	ReceivedMsgs <- msg

	helper.WriteJSON(w, http.StatusCreated, createResponse{msg.Id})
}

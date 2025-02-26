package controller

import (
	"chat-api/helper"
	"chat-api/model"
	"net/http"

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
	_, err = msg.Create()
	if err != nil {
		helper.ErrorJSON(w, err, http.StatusInternalServerError)
		return
	}

	receivedMsgs <- msg

	helper.WriteJSON(w, http.StatusCreated, createResponse{msg.Id})
}

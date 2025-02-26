package controller

import (
	"chat-api/helper"
	"chat-api/model"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func PostRoom(w http.ResponseWriter, r *http.Request) {
	room := model.Room{}
	err := helper.ReadJSON(r, &room)

	if err != nil {
		helper.ErrorJSON(w, err, http.StatusUnprocessableEntity)
		return
	}

	_, err = room.Create()
	if err != nil {
		helper.ErrorJSON(w, err, http.StatusInternalServerError)
		return
	}

	helper.WriteJSON(w, http.StatusCreated, createResponse{room.Id})
}

func GetRoom(w http.ResponseWriter, r *http.Request) {
	room := model.RoomItem{}

	id := chi.URLParam(r, "id")

	err := room.Get(id)

	if err != nil {
		helper.ErrorJSON(w, err, 500)
		return
	}

	helper.WriteJSON(w, http.StatusOK, room)
}

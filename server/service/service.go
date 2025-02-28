package service

import (
	"bytes"
	"chat-api/model"
	"encoding/json"
	"errors"
	"net/http"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type partialMessage struct {
	RoomId   primitive.ObjectID `json:"roomId"`
	Username string             `json:"username"`
	Text     string             `json:"text"`
}

func SubmitCommand(msg model.Message) error {
	pm := partialMessage{
		RoomId:   msg.RoomId,
		Username: msg.Username,
		Text:     msg.Text,
	}
	data, _ := json.Marshal(pm)

	request, err := http.NewRequest("POST", "http://localhost:8081/command", bytes.NewBuffer(data))
	if err != nil {
		return err
	}
	request.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return errors.New("invalid command")
	}

	return nil
}

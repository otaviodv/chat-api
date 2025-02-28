package test

import (
	"bytes"
	"chat-api/helper"
	"encoding/json"
	"errors"
	"net/http/httptest"
	"testing"
)

type testData struct {
	Message string `json:"message"`
}

func TestReadJSON(t *testing.T) {
	original := testData{Message: "new message"}
	jData, _ := json.Marshal(original)
	newData := testData{}

	req := httptest.NewRequest("GET", "http://site.com/hi", bytes.NewReader(jData))

	err := helper.ReadJSON(req, &newData)
	if err != nil {
		t.Error("unexpected error", err)
	}

	if newData.Message != original.Message {
		t.Errorf("expected value '%s', got '%s'", original.Message, newData.Message)
	}
}

func TestWriteJSON(t *testing.T) {
	original := testData{Message: "new message"}
	newData := testData{}

	w := httptest.NewRecorder()

	err := helper.WriteJSON(w, 200, original)

	if err != nil {
		t.Error("unexpected error", err)
	}

	resp := w.Result()
	err = json.NewDecoder(resp.Body).Decode(&newData)

	if err != nil {
		t.Error("unexpected error", err)
	}

	if newData.Message != original.Message {
		t.Errorf("expected value '%s', got '%s'", original.Message, newData.Message)
	}
}

func TestErrorJSON(t *testing.T) {
	type jsonResponse struct {
		Error   bool   `json:"error"`
		Message string `json:"message"`
		Data    any    `json:"data,omitempty"`
	}

	testError := errors.New("test error")

	w := httptest.NewRecorder()

	err := helper.ErrorJSON(w, testError)

	if err != nil {
		t.Error("unexpected error", err)
	}

	resp := w.Result()
	empty := jsonResponse{}
	err = json.NewDecoder(resp.Body).Decode(&empty)

	if err != nil {
		t.Error("unexpected error", err)
	}

	if testError.Error() != empty.Message {
		t.Errorf("expected value '%s', got '%s'", testError.Error(), empty.Message)
	}
}

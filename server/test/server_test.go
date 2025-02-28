package test

import (
	"bytes"
	"chat-api/helper"
	"encoding/json"
	"net/http/httptest"
	"testing"
)

type testData struct {
	Message string `json:"message"`
}

func TestReadJSON(t *testing.T) {
	original := testData{Message: "new message"}
	jData, _ := json.Marshal(original)
	empty := testData{}

	req := httptest.NewRequest("GET", "http://site.com/hi", bytes.NewReader(jData))

	err := helper.ReadJSON(req, &empty)
	if err != nil {
		t.Error("unexpected error")
	}

	if empty.Message != original.Message {
		t.Errorf("expected value '%s', got '%s'", empty.Message, original.Message)
	}
}

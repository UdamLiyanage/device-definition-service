package main

import (
	"encoding/json"
	"net/http"
	"testing"
)

func TestCreateDefinition(t *testing.T) {
	definition := DeviceDefinition{
		Version: "V1",
		Name:    "Test Case Device Definition",
		Type:    "Test Case Definition",
		Events:  []string{"publish", "subscribe"},
	}
	body, err := json.Marshal(definition)
	if err != nil {
		t.Fatal(err)
	}
	w := testRequestStatusCode("POST", "/device-definitions", body, http.StatusCreated, t)

	testDeleteDefinition(t, *w, &definition)
}

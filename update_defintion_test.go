package main

import (
	"encoding/json"
	"net/http"
	"testing"
)

func TestUpdateDefinition(t *testing.T) {
	var actions []map[string]string
	definition := DeviceDefinition{
		Version:    "v1",
		Name:       "New Definition Name",
		Type:       "Test Type",
		Events:     []string{"publish", "subscribe"},
		Actions:    append(actions, map[string]string{"type": "email"}),
		Parameters: map[string]string{"param1": "Param 1", "param2": "Param2"},
	}
	body, err := json.Marshal(definition)
	if err != nil {
		t.Fatal(err)
	}
	w := testRequestStatusCode("PUT", "/device-definitions/5e1dbd226f1cc1bed667e18e", body, http.StatusOK, t)
	testRequestBody(w, "MatchedCount", 0, t)
}

func TestUpdateDefinitionInvalid(t *testing.T) {
	var actions []map[string]string
	definition := DeviceDefinition{
		Version:    "v1",
		Name:       "New Definition Name",
		Type:       "Test Type",
		Events:     []string{"publish", "subscribe"},
		Actions:    append(actions, map[string]string{"type": "email"}),
		Parameters: map[string]string{"param1": "Param 1", "param2": "Param2"},
	}
	body, err := json.Marshal(definition)
	if err != nil {
		t.Fatal(err)
	}
	w := testRequestStatusCode("PUT", "/device-definitions/000000000000000000000000", body, http.StatusOK, t)
	testRequestBody(w, "MatchedCount", 0, t)
}

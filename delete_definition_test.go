package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func testDeleteDefinition(t *testing.T, w httptest.ResponseRecorder, definition *DeviceDefinition) {
	err := json.NewDecoder(w.Body).Decode(&definition)
	if err != nil {
		t.Fatal(err)
	}
	testRequestStatusCode("DELETE", "/device-definitions/"+definition.ID.Hex(), nil, http.StatusNotFound, t)
}

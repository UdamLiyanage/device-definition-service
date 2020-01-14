package main

import (
	"net/http"
	"testing"
)

func TestReadDefinition(t *testing.T) {
	//Object ID - 5e1d8c24bc2ee54c211f6f8e is always available as a test document
	testRequestStatusCode("GET", "/device-definitions/5e1d8c24bc2ee54c211f6f8e", nil, http.StatusOK, t)
}

func TestInvalidDeviceGet(t *testing.T) {
	testRequestStatusCode("GET", "/device-definitions/000000000000000000000000", nil, http.StatusNotFound, t)
}

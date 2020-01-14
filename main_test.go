package main

import "testing"

func TestServerPing(t *testing.T) {
	testRequestStatusCode("GET", "/ping", nil, 200, t)
}

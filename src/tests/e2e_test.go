package server_test

import (
	"testing"

	"net/http"
)

func TestServer(t *testing.T) {
	resp, error := http.Get("http://localhost:8090")

	if error != nil || resp.StatusCode != http.StatusOK {
		t.Errorf("server is not working properly")
	}
}

package server_test

import (
	"testing"

	server "github.com/perepadial/mdas-mds/src"
)

func TestWebContent(t *testing.T) {
	result := server.MainPageContent()
	if result == "" {
		t.Errorf("the function didn't return any content")
	}
}

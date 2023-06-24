package services

import (
	"net/http"
	"testing"
)

func TestRunService(t *testing.T) {
	go RunService()

	if _, err := http.Get("http://localhost:8080"); err != nil {
		t.Errorf("failed to start the service: %v", err)
	}
}

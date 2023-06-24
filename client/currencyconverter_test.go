package main

import (
	"io/ioutil"
	"net/http"
	"testing"
)

func TestClientServer(t *testing.T) {
	resp, err := http.Get("http://localhost:8000/")
	if err != nil {
		t.Fatalf("Failed to make GET request to client server: %v", err)
	}
	defer resp.Body.Close()

	_, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("Failed to read response body from client server: %v", err)
	}
}

func TestInternalServer(t *testing.T) {
	resp, err := http.Get("http://localhost:8080/")
	if err != nil {
		t.Fatalf("Failed to make GET request to internal server: %v", err)
	}
	defer resp.Body.Close()

	_, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("Failed to read response body from internal server: %v", err)
	}
}

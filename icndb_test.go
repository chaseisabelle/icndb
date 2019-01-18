package icndb

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPrepNames_success_timeSpace(t *testing.T) {
	first := "  chase"
	last := "isabelle   "

	names := prepNames(first, last)

	first = names["firstName"]
	last = names["lastName"]

	if first != "chase" {
		t.Errorf("Expected chase, but got %+v.", first)
	}

	if last != "isabelle" {
		t.Errorf("Expected isabelle, but got %+v.", last)
	}
}

func TestGet_failure_apiError(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
		response.WriteHeader(http.StatusOK)
		response.Header().Set("Content-Type", "application/json; charset=utf-8")
		response.Write([]byte(`{"type":"error","value":"you done messed up"}`))
	}))

	defer server.Close()

	host = server.URL

	_, err := get("", map[string]string{})

	if err == nil {
		t.Error("Expected error.")

		return
	}

	error := err.Error()

	if error != "error: you done messed up" {
		t.Errorf("Expected error: you done messed up, but got %+v.", error)
	}
}

func TestGet_failure_serverError(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
		response.WriteHeader(http.StatusInternalServerError)
		response.Header().Set("Content-Type", "application/json; charset=utf-8")
		response.Write([]byte(`Internal Server Error`))
	}))

	defer server.Close()

	host = server.URL

	_, err := get("", map[string]string{})

	if err == nil {
		t.Error("Expected error.")

		return
	}

	actual := err.Error()
	expected := http.StatusText(http.StatusInternalServerError)

	if actual != expected {
		t.Errorf("Expected %s, but got %+v.", expected, actual)
	}
}

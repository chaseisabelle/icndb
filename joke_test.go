package icndb

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetJoke_success_noNames(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
		response.WriteHeader(http.StatusOK)
		response.Header().Set("Content-Type", "application/json; charset=utf-8")
		response.Write([]byte(`{"type":"success","value":{"id":420,"joke":"lol","categories":[]}}`))
	}))

	defer server.Close()

	host = server.URL

	joke, err := GetJoke(420, "", "")

	if err != nil {
		t.Error(err)

		return
	}

	if joke.Id != 420 {
		t.Errorf("Expected 420, but got %+v.", joke.Id)
	}

	if joke.Text != "lol" {
		t.Errorf("Expected lol, but got %+v.", joke.Text)
	}
}

func TestGetJoke_success_firstName(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
		first := request.URL.Query().Get("firstName")

		if first != "chase" {
			t.Errorf("Expected chase, but got %+v.", first)
		}

		response.WriteHeader(http.StatusOK)
		response.Header().Set("Content-Type", "application/json; charset=utf-8")
		response.Write([]byte(`{"type":"success","value":{"id":420,"joke":"lol","categories":[]}}`))
	}))

	defer server.Close()

	host = server.URL

	_, err := GetJoke(420, "chase", "")

	if err != nil {
		t.Error(err)
	}
}

func TestGetJoke_success_lastName(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
		last := request.URL.Query().Get("lastName")

		if last != "isabelle" {
			t.Errorf("Expected isabelle, but got %+v.", last)
		}

		response.WriteHeader(http.StatusOK)
		response.Header().Set("Content-Type", "application/json; charset=utf-8")
		response.Write([]byte(`{"type":"success","value":{"id":420,"joke":"lol","categories":[]}}`))
	}))

	defer server.Close()

	host = server.URL

	_, err := GetJoke(420, "", "isabelle")

	if err != nil {
		t.Error(err)
	}
}

func TestGetJoke_success_bothNames(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
		first := request.URL.Query().Get("firstName")

		if first != "chase" {
			t.Errorf("Expected chase, but got %+v.", first)
		}

		last := request.URL.Query().Get("lastName")

		if last != "isabelle" {
			t.Errorf("Expected isabelle, but got %+v.", last)
		}

		response.WriteHeader(http.StatusOK)
		response.Header().Set("Content-Type", "application/json; charset=utf-8")
		response.Write([]byte(`{"type":"success","value":{"id":420,"joke":"lol","categories":[]}}`))
	}))

	defer server.Close()

	host = server.URL

	_, err := GetJoke(420, "chase", "isabelle")

	if err != nil {
		t.Error(err)
	}
}

func TestGetRandomJoke_success(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
		response.WriteHeader(http.StatusOK)
		response.Header().Set("Content-Type", "application/json; charset=utf-8")
		response.Write([]byte(`{"type":"success","value":{"id":666,"joke":"random","categories":[]}}`))
	}))

	defer server.Close()

	host = server.URL

	joke, err := GetRandomJoke("", "")

	if err != nil {
		t.Error(err)

		return
	}

	if joke.Id != 666 {
		t.Errorf("Expected 666, but got %+v.", joke.Id)
	}

	if joke.Text != "random" {
		t.Errorf("Expected random, but got %+v.", joke.Text)
	}
}

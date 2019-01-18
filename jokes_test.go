package icndb

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetJokes_success_noNames(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
		response.WriteHeader(http.StatusOK)
		response.Header().Set("Content-Type", "application/json; charset=utf-8")
		response.Write([]byte(`{"type":"success","value":[{"id":69,"joke":"lol","categories":[]},{"id":666,"joke":"wtf","categories":["turtles"]}]}`))
	}))

	defer server.Close()

	host = server.URL

	jokes, err := GetJokes("", "")

	if err != nil {
		t.Error(err)

		return
	}

	count := len(jokes)

	if count != 2 {
		t.Errorf("Expected 2, but got %+v.", count)

		return
	}

	joke := jokes[0]

	if joke.Id != 69 {
		t.Errorf("Expected 69, but got %+v.", joke.Id)
	}

	if joke.Text != "lol" {
		t.Errorf("Expected lol, but got %+v.", joke.Text)
	}

	joke = jokes[1]

	if joke.Id != 666 {
		t.Errorf("Expected 666, but got %+v.", joke.Id)
	}

	if joke.Text != "wtf" {
		t.Errorf("Expected wtf, but got %+v.", joke.Text)
	}
}

func TestGetJokes_success_firstName(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
		first := request.URL.Query().Get("firstName")

		if first != "chase" {
			t.Errorf("Expected chase, but got %+v.", first)
		}

		response.WriteHeader(http.StatusOK)
		response.Header().Set("Content-Type", "application/json; charset=utf-8")
		response.Write([]byte(`{"type":"success","value":[{"id":69,"joke":"lol","categories":[]}]}`))
	}))

	defer server.Close()

	host = server.URL

	_, err := GetJokes("chase", "")

	if err != nil {
		t.Error(err)
	}
}

func TestGetJokes_success_lastName(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
		last := request.URL.Query().Get("lastName")

		if last != "isabelle" {
			t.Errorf("Expected isabelle, but got %+v.", last)
		}

		response.WriteHeader(http.StatusOK)
		response.Header().Set("Content-Type", "application/json; charset=utf-8")
		response.Write([]byte(`{"type":"success","value":[{"id":69,"joke":"lol","categories":[]}]}`))
	}))

	defer server.Close()

	host = server.URL

	_, err := GetJokes("", "isabelle")

	if err != nil {
		t.Error(err)
	}
}

func TestGetJokes_success_bothNames(t *testing.T) {
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
		response.Write([]byte(`{"type":"success","value":[{"id":69,"joke":"lol","categories":[]}]}`))
	}))

	defer server.Close()

	host = server.URL

	_, err := GetJokes("chase", "isabelle")

	if err != nil {
		t.Error(err)
	}
}

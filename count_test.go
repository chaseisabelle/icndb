package icndb

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetCount_success(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
		response.WriteHeader(http.StatusOK)
		response.Header().Set("Content-Type", "application/json; charset=utf-8")
		response.Write([]byte(`{"type":"success","value":420}`))
	}))

	defer server.Close()

	count, err := icndb(server.URL).Count()

	if err != nil {
		t.Error(err)

		return
	}

	if count != 420 {
		t.Errorf("Expected 420, but got %+v.", count)
	}
}

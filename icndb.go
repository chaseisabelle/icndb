package icndb

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strings"
)

const HOST = "http://api.icndb.com"

type ICNDB struct {
	Client *http.Client
	Host   string
}

func New() *ICNDB {
	return &ICNDB{
		Client: &http.Client{},
		Host:   HOST,
	}
}

func (icndb *ICNDB) get(endpoint string, params map[string]string) (interface{}, error) {
	request, err := http.NewRequest("GET", icndb.Host+"/"+endpoint, nil)

	if err != nil {
		return nil, err
	}

	query := request.URL.Query()

	for name, param := range params {
		query.Add(name, param)
	}

	request.URL.RawQuery = query.Encode()
	response, err := icndb.Client.Do(request)

	if response != nil {
		defer response.Body.Close()
	}

	if err != nil {
		return nil, err
	}

	if response.StatusCode != http.StatusOK {
		return nil, errors.New(http.StatusText(response.StatusCode))
	}

	body, err := ioutil.ReadAll(response.Body)

	if err != nil {
		return nil, err
	}

	var payload map[string]interface{}

	err = json.Unmarshal(body, &payload)

	if err != nil {
		return nil, err
	}

	status, ok := payload["type"].(string)

	if !ok {
		return nil, errors.New("Failed to detect response status.")
	}

	value := payload["value"]

	if status == "success" {
		return value, nil
	}

	error, ok := value.(string)

	if !ok || error == "" {
		error = "Failed to detect error message."
	}

	return nil, errors.New(status + ": " + error)
}

func prepNames(first string, last string) map[string]string {
	first = strings.TrimSpace(first)
	last = strings.TrimSpace(last)

	names := make(map[string]string)

	if first != "" {
		names["firstName"] = first
	}

	if last != "" {
		names["lastName"] = last
	}

	return names
}

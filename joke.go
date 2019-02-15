package icndb

import (
	"errors"
	"fmt"
)

type Joke struct {
	Id   uint64 `json:"id"`
	Text string `json:"text"`
}

func (icndb *ICNDB) Joke(id uint64, first string, last string) (*Joke, error) {
	payload, err := icndb.get(fmt.Sprintf("jokes/%d", id), prepNames(first, last))

	if err != nil {
		return nil, err
	}

	return buildJoke(payload)
}

func (icndb *ICNDB) RandomJoke(first string, last string) (*Joke, error) {
	payload, err := icndb.get("jokes/random", prepNames(first, last))

	if err != nil {
		return nil, err
	}

	return buildJoke(payload)
}

func buildJoke(payload interface{}) (*Joke, error) {
	mapped, ok := payload.(map[string]interface{})

	if !ok {
		return nil, errors.New("Failed to map joke interface.")
	}

	id, ok := mapped["id"].(float64)

	if !ok {
		return nil, errors.New("Failed to get joke ID.")
	}

	text, ok := mapped["joke"].(string)

	if !ok {
		return nil, errors.New("Failed to get joke text.")
	}

	return &Joke{
		Id:   uint64(id),
		Text: text,
	}, nil
}

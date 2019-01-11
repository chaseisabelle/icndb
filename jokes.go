package icndb

import (
	"errors"
	"fmt"
)

func GetJokes(first string, last string) ([]*Joke, error) {
	payload, err := get("jokes", prepNames(first, last))

	if err != nil {
		return nil, err
	}

	values, ok := payload.([]interface{})

	if !ok {
		return nil, errors.New("Failed to load jokes.")
	}

	return buildJokes(values)
}

func GetRandomJokes(count uint64, first string, last string) ([]*Joke, error) {
	payload, err := get(fmt.Sprintf("jokes/random/%d", count), prepNames(first, last))

	if err != nil {
		return nil, err
	}

	values, ok := payload.([]interface{})

	if !ok {
		return nil, errors.New("Failed to load jokes.")
	}

	return buildJokes(values)
}

func buildJokes(payloads []interface{}) ([]*Joke, error) {
	var jokes []*Joke

	for _, payload := range payloads {
		joke, err := buildJoke(payload)

		if err != nil {
			return nil, err
		}

		jokes = append(jokes, joke)
	}

	return jokes, nil
}

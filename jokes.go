package icndb

import (
	"errors"
	"fmt"
	"strings"
)

func (icndb *ICNDB) Jokes(first string, last string) ([]*Joke, error) {
	payload, err := icndb.get("jokes", prepNames(first, last))

	if err != nil {
		return nil, err
	}

	values, ok := payload.([]interface{})

	if !ok {
		return nil, errors.New("Failed to load jokes.")
	}

	return buildJokes(values)
}

func (icndb *ICNDB) RandomJokes(count uint64, first string, last string, categories map[string]bool) ([]*Joke, error) {
	inclusions := []string{}
	exclusions := []string{}

	for category, include := range categories {
		if include {
			inclusions = append(inclusions, category)
		} else {
			exclusions = append(exclusions, category)
		}
	}

	params := prepNames(first, last)

	if len(inclusions) > 0 {
		params["limitTo"] = "[" + strings.Join(inclusions, ",") + "]"
	}

	if len(exclusions) > 0 {
		params["exclude"] = "[" + strings.Join(inclusions, ",") + "]"
	}

	payload, err := icndb.get(fmt.Sprintf("jokes/random/%d", count), params)

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

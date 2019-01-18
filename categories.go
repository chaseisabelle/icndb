package icndb

import "github.com/pkg/errors"

func GetCategories() ([]string, error) {
	payload, err := get("categories", make(map[string]string))

	if err != nil {
		return nil, err
	}

	values, ok := payload.([]interface{})

	if !ok {
		return nil, errors.New("Failed to convert interface to array.")
	}

	var categories []string

	for _, value := range values {
		category, ok := value.(string)

		if !ok {
			return nil, errors.New("Failed to convert interface to string.")
		}

		categories = append(categories, category)
	}

	return categories, nil
}

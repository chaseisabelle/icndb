package icndb

import "errors"

func GetCount() (uint64, error) {
	payload, err := get("count", make(map[string]string))

	if err != nil {
		return 0, err
	}

	count, ok := payload.(float64)

	if !ok {
		return 0, errors.New("Failed to convert interface to float.")
	}

	return uint64(count), nil
}

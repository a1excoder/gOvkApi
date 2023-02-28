package gOvkApi

import "encoding/json"

func isError(data []byte) (bool, error) {
	value := make(map[string]any)

	if err := json.Unmarshal(data, &value); err != nil {
		return false, err
	}

	if _, found := value["error_code"]; found {
		return true, nil
	}

	return false, nil
}

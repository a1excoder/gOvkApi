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

func unmarshalAny[T any](bytes []byte) (*T, error) {
	out := new(T)
	if err := json.Unmarshal(bytes, out); err != nil {
		return nil, err
	}
	return out, nil
}

func unmarshalError(data []byte) (*ErrorReturned, error) {
	return unmarshalAny[ErrorReturned](data)
}

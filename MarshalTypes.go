package gOvkApi

import "encoding/json"

func isError(data []byte) (*ovkError, error) {
	errPtr, err := unmarshalError(data)
	if err != nil {
		return nil, err
	}

	if errPtr.ErrorMsg != "" {
		return errPtr, nil
	}

	return nil, nil
}

func unmarshalAny[T any](bytes []byte) (*T, error) {
	out := new(T)
	if err := json.Unmarshal(bytes, out); err != nil {
		return nil, err
	}
	return out, nil
}

func unmarshalError(data []byte) (*ovkError, error) {
	return unmarshalAny[ovkError](data)
}

func unmarshalSuccessToken(data []byte) (*SuccessToken, error) {
	return unmarshalAny[SuccessToken](data)
}

func GetAuthData(instance string) *AuthData {
	return &AuthData{Instance: instance}
}

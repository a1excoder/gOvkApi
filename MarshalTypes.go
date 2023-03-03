package gOvkApi

import "encoding/json"

func isError(data []byte) (*ErrorReturned, bool, error) {
	errPtr, err := unmarshalError(data)
	if err != nil {
		return nil, false, err
	}

	if errPtr.ErrorMsg != "" {
		return errPtr, true, nil
	}

	return nil, false, nil
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

func unmarshalSuccessToken(data []byte) (*SuccessToken, error) {
	return unmarshalAny[SuccessToken](data)
}

func getAuthData(instance string) *AuthData {
	authPtr := new(AuthData)
	authPtr.Instance = instance

	return authPtr
}

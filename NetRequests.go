package gOvkApi

import (
	"io"
	"net/http"
)

const (
	MethodPost = 1
	MethodGet  = 2
)

func makeRequest(url string, method int) ([]byte, int, error) {
	var response *http.Response
	var err error

	switch method {
	case MethodPost:
		response, err = http.Post(url, "multipart/form-data", nil)
		break
	case MethodGet:
		response, err = http.Get(url)
		break
	default:
		response, err = http.Post(url, "multipart/form-data", nil)
	}

	if err != nil {
		return nil, 0, err
	}

	bodyData, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, 0, err
	}

	return bodyData, response.StatusCode, nil
}

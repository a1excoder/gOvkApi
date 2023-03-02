package gOvkApi

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

const (
	MethodPost = 1
	MethodGet  = 2
)

func makeRequest(url string, params url.Values, method int) ([]byte, int, error) {
	var response *http.Response
	var err error

	switch method {
	case MethodPost:
		//response, err = http.Post(url, "multipart/form-data", nil)
		response, err = http.PostForm(url, params)
		break
	case MethodGet:
		response, err = http.Get(url + params.Encode())
		break
	default:
		response, err = http.PostForm(url, params)
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

func (data *AuthData) getApiToken(username, password string) (*ErrorReturned, error) {
	body, _, err := makeRequest(fmt.Sprintf("%s/token?username=%s&password=%s&grant_type=password",
		data.Instance, username, password), MethodGet)
	if err != nil {
		return nil, err
	}

	ovkErr, isErr, err := isError(body)
	if err != nil {
		return nil, err
	}

	if isErr {
		return ovkErr, nil
	}

	data.Token, err = unmarshalSuccessToken(body)
	if err != nil {
		return nil, err
	}

	return nil, nil
}

package gOvkApi

import (
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
		response, err = http.PostForm(url, params)
		break
	case MethodGet:
		response, err = http.Get(url + "?" + params.Encode())
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

func (authData *AuthData) GetApiToken(username, password, twoFactorCode string, code bool) (*ErrorReturned, error) {
	params := url.Values{}
	params.Add("username", username)
	params.Add("password", password)
	params.Add("grant_type", "password")

	if code {
		params.Add("code", twoFactorCode)
	}

	body, _, err := makeRequest(authData.Instance+"/token", params, MethodPost)
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

	authData.Token, err = unmarshalSuccessToken(body)
	if err != nil {
		return nil, err
	}

	return nil, nil
}

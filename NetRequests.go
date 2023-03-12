package gOvkApi

import (
	"io"
	"net/http"
	"net/url"
)

const (
	methodPost = 1
	methodGet  = 2
)

func makeRequest(url string, params url.Values, method int) ([]byte, int, error) {
	var response *http.Response
	var err error

	switch method {
	case methodPost:
		response, err = http.PostForm(url, params)
		break
	case methodGet:
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

	_ = response.Body.Close()
	return bodyData, response.StatusCode, nil
}

func (authData *AuthData) GetApiToken(username, password, twoFactorCode string, code bool) ErrorReturned {
	params := url.Values{}
	params.Add("username", username)
	params.Add("password", password)
	params.Add("grant_type", "password")

	if code {
		params.Add("code", twoFactorCode)
	}

	errRet := ErrorReturned{}

	body, _, err := makeRequest(authData.Instance+"/token", params, methodPost)
	if err != nil {
		errRet.Err = err
		return errRet
	}

	errRet.OvkError, errRet.Err = isError(body)
	if errRet.Err != nil || errRet.OvkError != nil {
		return errRet
	}

	authData.Token, err = unmarshalSuccessToken(body)
	if err != nil {
		errRet.Err = err
		return errRet
	}

	return errRet
}

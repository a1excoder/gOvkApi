package gOvkApi

import (
	"net/url"
	"strconv"
)

func (authData *AuthData) getFriends(userId, count, offset int) (*GetFriends, *ErrorReturned, error) {
	params := url.Values{}
	params.Add("access_token", authData.Token.AccessToken)
	params.Add("user_id", strconv.Itoa(userId))
	params.Add("count", strconv.Itoa(count))
	params.Add("offset", strconv.Itoa(offset))

	body, _, err := makeRequest(authData.Instance+"/method/Friends.get", params, MethodGet)
	if err != nil {
		return nil, nil, err
	}

	ovkErr, isErr, err := isError(body)
	if err != nil {
		return nil, nil, err
	}

	if isErr {
		return nil, ovkErr, nil
	}

	profileInfo, err := unmarshalAny[GetFriends](body)
	if err != nil {
		return nil, nil, err
	}

	return profileInfo, nil, nil
}

func (authData *AuthData) addFriend(userId int) (*ResponseFriends, *ErrorReturned, error) {
	params := url.Values{}
	params.Add("access_token", authData.Token.AccessToken)
	params.Add("user_id", strconv.Itoa(userId))

	body, _, err := makeRequest(authData.Instance+"/method/Friends.add", params, MethodGet)
	if err != nil {
		return nil, nil, err
	}

	ovkErr, isErr, err := isError(body)
	if err != nil {
		return nil, nil, err
	}

	if isErr {
		return nil, ovkErr, nil
	}

	respFriends, err := unmarshalAny[ResponseFriends](body)
	if err != nil {
		return nil, nil, err
	}

	return respFriends, nil, nil
}

func (authData *AuthData) deleteFriend(userId int) (*ResponseFriends, *ErrorReturned, error) {
	params := url.Values{}
	params.Add("access_token", authData.Token.AccessToken)
	params.Add("user_id", strconv.Itoa(userId))

	body, _, err := makeRequest(authData.Instance+"/method/Friends.delete", params, MethodGet)
	if err != nil {
		return nil, nil, err
	}

	ovkErr, isErr, err := isError(body)
	if err != nil {
		return nil, nil, err
	}

	if isErr {
		return nil, ovkErr, nil
	}

	respFriends, err := unmarshalAny[ResponseFriends](body)
	if err != nil {
		return nil, nil, err
	}

	return respFriends, nil, nil
}

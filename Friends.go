package gOvkApi

import (
	"net/url"
	"strconv"
)

func (authData *AuthData) GetFriends(userId, count, offset int) (*GetFriends, *ErrorReturned, error) {
	params := url.Values{}
	params.Add("access_token", authData.Token.AccessToken)
	params.Add("user_id", strconv.Itoa(userId))
	params.Add("count", strconv.Itoa(count))
	params.Add("offset", strconv.Itoa(offset))

	body, _, err := makeRequest(authData.Instance+"/method/Friends.get", params, MethodGet)
	if err != nil {
		return nil, nil, err
	}

	ovkErr, err := isError(body)
	if err != nil {
		return nil, nil, err
	}

	if ovkErr != nil {
		return nil, ovkErr, nil
	}

	getFriends, err := unmarshalAny[GetFriends](body)
	if err != nil {
		return nil, nil, err
	}

	return getFriends, nil, nil
}

func (authData *AuthData) AddFriend(userId int) (*ResponseFriends, *ErrorReturned, error) {
	params := url.Values{}
	params.Add("access_token", authData.Token.AccessToken)
	params.Add("user_id", strconv.Itoa(userId))

	body, _, err := makeRequest(authData.Instance+"/method/Friends.add", params, MethodGet)
	if err != nil {
		return nil, nil, err
	}

	ovkErr, err := isError(body)
	if err != nil {
		return nil, nil, err
	}

	if ovkErr != nil {
		return nil, ovkErr, nil
	}

	respFriends, err := unmarshalAny[ResponseFriends](body)
	if err != nil {
		return nil, nil, err
	}

	return respFriends, nil, nil
}

func (authData *AuthData) DeleteFriend(userId int) (*ResponseFriends, *ErrorReturned, error) {
	params := url.Values{}
	params.Add("access_token", authData.Token.AccessToken)
	params.Add("user_id", strconv.Itoa(userId))

	body, _, err := makeRequest(authData.Instance+"/method/Friends.delete", params, MethodGet)
	if err != nil {
		return nil, nil, err
	}

	ovkErr, err := isError(body)
	if err != nil {
		return nil, nil, err
	}

	if ovkErr != nil {
		return nil, ovkErr, nil
	}

	respFriends, err := unmarshalAny[ResponseFriends](body)
	if err != nil {
		return nil, nil, err
	}

	return respFriends, nil, nil
}

func (authData *AuthData) GetRequests(extended, count, offset int) (*GetRequests, *ErrorReturned, error) {
	params := url.Values{}
	params.Add("access_token", authData.Token.AccessToken)
	params.Add("extended", strconv.Itoa(extended))
	params.Add("count", strconv.Itoa(count))
	params.Add("offset", strconv.Itoa(offset))

	body, _, err := makeRequest(authData.Instance+"/method/Friends.getRequests", params, MethodGet)
	if err != nil {
		return nil, nil, err
	}

	ovkErr, err := isError(body)
	if err != nil {
		return nil, nil, err
	}

	if ovkErr != nil {
		return nil, ovkErr, nil
	}

	getReqFriends, err := unmarshalAny[GetRequests](body)
	if err != nil {
		return nil, nil, err
	}

	return getReqFriends, nil, nil
}

func (authData *AuthData) AreFriends(userIds string) (*AreFriends, *ErrorReturned, error) {
	params := url.Values{}
	params.Add("access_token", authData.Token.AccessToken)
	params.Add("user_ids", userIds)

	body, _, err := makeRequest(authData.Instance+"/method/Friends.areFriends", params, MethodGet)
	if err != nil {
		return nil, nil, err
	}

	ovkErr, err := isError(body)
	if err != nil {
		return nil, nil, err
	}

	if ovkErr != nil {
		return nil, ovkErr, nil
	}

	areFriendsReq, err := unmarshalAny[AreFriends](body)
	if err != nil {
		return nil, nil, err
	}

	return areFriendsReq, nil, nil
}

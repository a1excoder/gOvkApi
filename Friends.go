package gOvkApi

import (
	"net/url"
	"strconv"
)

func (authData *AuthData) GetFriends(userId, count, offset int) (*GetFriends, ErrorReturned) {
	params := url.Values{}
	params.Add("access_token", authData.Token.AccessToken)
	params.Add("user_id", strconv.Itoa(userId))
	params.Add("count", strconv.Itoa(count))
	params.Add("offset", strconv.Itoa(offset))

	errRet := ErrorReturned{}

	body, _, err := makeRequest(authData.Instance+"/method/Friends.get", params, methodGet)
	if err != nil {
		errRet.Err = err
		return nil, errRet
	}

	errRet.OvkError, errRet.Err = isError(body)
	if errRet.Err != nil || errRet.OvkError != nil {
		return nil, errRet
	}

	getFriends, err := unmarshalAny[GetFriends](body)
	if err != nil {
		errRet.Err = err
		return nil, errRet
	}

	return getFriends, errRet
}

func (authData *AuthData) AddFriend(userId int) (*ResponseFriends, ErrorReturned) {
	params := url.Values{}
	params.Add("access_token", authData.Token.AccessToken)
	params.Add("user_id", strconv.Itoa(userId))

	errRet := ErrorReturned{}

	body, _, err := makeRequest(authData.Instance+"/method/Friends.add", params, methodGet)
	if err != nil {
		errRet.Err = err
		return nil, errRet
	}

	errRet.OvkError, errRet.Err = isError(body)
	if errRet.Err != nil || errRet.OvkError != nil {
		return nil, errRet
	}

	respFriends, err := unmarshalAny[ResponseFriends](body)
	if err != nil {
		errRet.Err = err
		return nil, errRet
	}

	return respFriends, errRet
}

func (authData *AuthData) DeleteFriend(userId int) (*ResponseFriends, ErrorReturned) {
	params := url.Values{}
	params.Add("access_token", authData.Token.AccessToken)
	params.Add("user_id", strconv.Itoa(userId))

	errRet := ErrorReturned{}

	body, _, err := makeRequest(authData.Instance+"/method/Friends.delete", params, methodGet)
	if err != nil {
		errRet.Err = err
		return nil, errRet
	}

	errRet.OvkError, errRet.Err = isError(body)
	if errRet.Err != nil || errRet.OvkError != nil {
		return nil, errRet
	}

	respFriends, err := unmarshalAny[ResponseFriends](body)
	if err != nil {
		errRet.Err = err
		return nil, errRet
	}

	return respFriends, errRet
}

func (authData *AuthData) GetRequests(extended, count, offset int) (*GetRequests, ErrorReturned) {
	params := url.Values{}
	params.Add("access_token", authData.Token.AccessToken)
	params.Add("extended", strconv.Itoa(extended))
	params.Add("count", strconv.Itoa(count))
	params.Add("offset", strconv.Itoa(offset))

	errRet := ErrorReturned{}

	body, _, err := makeRequest(authData.Instance+"/method/Friends.getRequests", params, methodGet)
	if err != nil {
		errRet.Err = err
		return nil, errRet
	}

	errRet.OvkError, errRet.Err = isError(body)
	if errRet.Err != nil || errRet.OvkError != nil {
		return nil, errRet
	}

	getReqFriends, err := unmarshalAny[GetRequests](body)
	if err != nil {
		errRet.Err = err
		return nil, errRet
	}

	return getReqFriends, errRet
}

func (authData *AuthData) AreFriends(userIds string) (*AreFriends, ErrorReturned) {
	params := url.Values{}
	params.Add("access_token", authData.Token.AccessToken)
	params.Add("user_ids", userIds)

	errRet := ErrorReturned{}

	body, _, err := makeRequest(authData.Instance+"/method/Friends.areFriends", params, methodGet)
	if err != nil {
		errRet.Err = err
		return nil, errRet
	}

	errRet.OvkError, errRet.Err = isError(body)
	if errRet.Err != nil || errRet.OvkError != nil {
		return nil, errRet
	}

	areFriendsReq, err := unmarshalAny[AreFriends](body)
	if err != nil {
		errRet.Err = err
		return nil, errRet
	}

	return areFriendsReq, errRet
}

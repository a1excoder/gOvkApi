package gOvkApi

import (
	"fmt"
	"net/url"
	"strconv"
	"strings"
)

func (authData *AuthData) GetGroups(userId, count, offset int, fields string) (*GetGroups, ErrorReturned) {
	params := url.Values{}
	params.Add("access_token", authData.Token.AccessToken)
	params.Add("user_id", strconv.Itoa(userId))
	params.Add("count", strconv.Itoa(count))
	params.Add("offset", strconv.Itoa(offset))
	params.Add("fields", fields)

	errRet := ErrorReturned{}

	body, _, err := makeRequest(authData.Instance+"/method/Groups.get", params, methodGet)
	if err != nil {
		errRet.Err = err
		return nil, errRet
	}

	errRet.OvkError, errRet.Err = isError(body)
	if errRet.Err != nil || errRet.OvkError != nil {
		return nil, errRet
	}

	getGroups, err := unmarshalAny[GetGroups](body)
	if err != nil {
		errRet.Err = err
		return nil, errRet
	}

	return getGroups, errRet
}

func (authData *AuthData) GetById(groupIds []int, count, offset int, fields string) (*GetById, ErrorReturned) {
	params := url.Values{}
	params.Add("access_token", authData.Token.AccessToken)
	params.Add("group_ids", strings.Trim(strings.Join(strings.Fields(fmt.Sprint(groupIds)), ","), "[]"))
	params.Add("count", strconv.Itoa(count))
	params.Add("offset", strconv.Itoa(offset))
	params.Add("fields", fields)

	errRet := ErrorReturned{}

	body, _, err := makeRequest(authData.Instance+"/method/Groups.getById", params, methodGet)
	if err != nil {
		errRet.Err = err
		return nil, errRet
	}

	errRet.OvkError, errRet.Err = isError(body)
	if errRet.Err != nil || errRet.OvkError != nil {
		return nil, errRet
	}

	getById, err := unmarshalAny[GetById](body)
	if err != nil {
		errRet.Err = err
		return nil, errRet
	}

	return getById, errRet
}

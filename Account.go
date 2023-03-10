package gOvkApi

import "net/url"

func (authData *AuthData) GetProfileInfo() (*ProfileInfo, ErrorReturned) {
	params := url.Values{}
	params.Add("access_token", authData.Token.AccessToken)

	errRet := ErrorReturned{}

	body, _, err := makeRequest(authData.Instance+"/method/Account.getProfileInfo", params, methodGet)
	if err != nil {
		errRet.Err = err
		return nil, errRet
	}

	errRet.OvkError, errRet.Err = isError(body)
	if errRet.Err != nil || errRet.OvkError != nil {
		return nil, errRet
	}

	profileInfo, err := unmarshalAny[ProfileInfo](body)
	if err != nil {
		errRet.Err = err
		return nil, errRet
	}

	return profileInfo, errRet
}

func (authData *AuthData) GetInfo() (*Info, ErrorReturned) {
	params := url.Values{}
	params.Add("access_token", authData.Token.AccessToken)

	errRet := ErrorReturned{}

	body, _, err := makeRequest(authData.Instance+"/method/Account.getInfo", params, methodGet)
	if err != nil {
		errRet.Err = err
		return nil, errRet
	}

	errRet.OvkError, errRet.Err = isError(body)
	if errRet.Err != nil || errRet.OvkError != nil {
		return nil, errRet
	}

	info, err := unmarshalAny[Info](body)
	if err != nil {
		errRet.Err = err
		return nil, errRet
	}

	return info, errRet
}

func (authData *AuthData) GetCounters() (*Counters, ErrorReturned) {
	params := url.Values{}
	params.Add("access_token", authData.Token.AccessToken)

	errRet := ErrorReturned{}

	body, _, err := makeRequest(authData.Instance+"/method/Account.getCounters", params, methodGet)
	if err != nil {
		errRet.Err = err
		return nil, errRet
	}

	errRet.OvkError, errRet.Err = isError(body)
	if errRet.Err != nil || errRet.OvkError != nil {
		return nil, errRet
	}

	counters, err := unmarshalAny[Counters](body)
	if err != nil {
		errRet.Err = err
		return nil, errRet
	}

	return counters, errRet
}

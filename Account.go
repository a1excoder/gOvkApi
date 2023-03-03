package gOvkApi

import "net/url"

func (authData *AuthData) getProfileInfo() (*ProfileInfo, *ErrorReturned, error) {
	params := url.Values{}
	params.Add("access_token", authData.Token.AccessToken)

	body, _, err := makeRequest(authData.Instance+"/method/Account.getProfileInfo", params, MethodGet)
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

	profileInfo, err := unmarshalAny[ProfileInfo](body)
	if err != nil {
		return nil, nil, err
	}

	return profileInfo, nil, nil
}

func (authData *AuthData) getInfo() (*Info, *ErrorReturned, error) {
	params := url.Values{}
	params.Add("access_token", authData.Token.AccessToken)

	body, _, err := makeRequest(authData.Instance+"/method/Account.getInfo", params, MethodGet)
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

	info, err := unmarshalAny[Info](body)
	if err != nil {
		return nil, nil, err
	}

	return info, nil, nil
}

func (authData *AuthData) getCounters() (*Counters, *ErrorReturned, error) {
	params := url.Values{}
	params.Add("access_token", authData.Token.AccessToken)

	body, _, err := makeRequest(authData.Instance+"/method/Account.getCounters", params, MethodGet)
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

	counters, err := unmarshalAny[Counters](body)
	if err != nil {
		return nil, nil, err
	}

	return counters, nil, nil
}

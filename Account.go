package gOvkApi

import "net/url"

func (authData *AuthData) GetProfileInfo() (*ProfileInfo, *ErrorReturned, error) {
	params := url.Values{}
	params.Add("access_token", authData.Token.AccessToken)

	body, _, err := makeRequest(authData.Instance+"/method/Account.getProfileInfo", params, methodGet)
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

	profileInfo, err := unmarshalAny[ProfileInfo](body)
	if err != nil {
		return nil, nil, err
	}

	return profileInfo, nil, nil
}

func (authData *AuthData) GetInfo() (*Info, *ErrorReturned, error) {
	params := url.Values{}
	params.Add("access_token", authData.Token.AccessToken)

	body, _, err := makeRequest(authData.Instance+"/method/Account.getInfo", params, methodGet)
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

	info, err := unmarshalAny[Info](body)
	if err != nil {
		return nil, nil, err
	}

	return info, nil, nil
}

func (authData *AuthData) GetCounters() (*Counters, *ErrorReturned, error) {
	params := url.Values{}
	params.Add("access_token", authData.Token.AccessToken)

	body, _, err := makeRequest(authData.Instance+"/method/Account.getCounters", params, methodGet)
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

	counters, err := unmarshalAny[Counters](body)
	if err != nil {
		return nil, nil, err
	}

	return counters, nil, nil
}

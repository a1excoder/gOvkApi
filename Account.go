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

package gOvkApi

type AuthData struct {
	Instance string
	Token    *SuccessToken
}

type SuccessToken struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
	UserId      int    `json:"user_id"`
}

type ErrorReturned struct {
	ErrorCode int    `json:"error_code"`
	ErrorMsg  string `json:"error_msg"`
	//	RequestParams []RequestParams[any] `json:"request_params"`
}

//type RequestParams[T any] struct {
//	Key   string `json:"key"`
//	Value T      `json:"value"`
//}

type ProfileInfo struct {
	Response struct {
		FirstName              string `json:"first_name"`
		Id                     int    `json:"id"`
		LastName               string `json:"last_name"`
		HomeTown               string `json:"home_town"`
		Status                 string `json:"status"`
		BirthdayDate           string `json:"bdate"`
		BirthdayDateVisibility int    `json:"bdate_visibility"`
		Phone                  string `json:"phone"`
		Relation               int    `json:"relation"`
		Sex                    int    `json:"sex"`
	} `json:"response"`
}

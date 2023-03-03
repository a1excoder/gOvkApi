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

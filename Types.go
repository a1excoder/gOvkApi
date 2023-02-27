package gOvkApi

type SuccessToken struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
	UserId      int    `json:"user_id"`
}

type ErrorReturned struct {
	ErrorCode    int                 `json:"error_code"`
	ErrorMsg     string              `json:"error_msg"`
	RequestParam []RequestParam[any] `json:"request_param"`
}

type RequestParam[T any] struct {
	Key   string `json:"key"`
	Value T      `json:"value"`
}

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

type Info struct {
	Response struct {
		FaRequired                int    `json:"2fa_required"`
		Country                   string `json:"country"`
		EuUser                    bool   `json:"eu_user"`
		HttpsRequired             int    `json:"https_required"`
		Intro                     int    `json:"intro"`
		CommunityComments         bool   `json:"community_comments"`
		IsLiveStreamingEnabled    bool   `json:"is_live_streaming_enabled"`
		IsNewLiveStreamingEnabled bool   `json:"is_new_live_streaming_enabled"`
		Lang                      int    `json:"lang"`
		NoWallReplies             int    `json:"no_wall_replies"`
		OwnPostsDefault           int    `json:"own_posts_default"`
	} `json:"response"`
}

type Counters struct {
	Response struct {
		Friends       int `json:"friends"`
		Notifications int `json:"notifications"`
		Messages      int `json:"messages"`
	} `json:"response"`
}

type GetFriends struct {
	Response struct {
		Count int `json:"count"`
		Items []struct {
			Id              int    `json:"id"`
			FirstName       string `json:"first_name"`
			LastName        string `json:"last_name"`
			IsClosed        bool   `json:"is_closed"`
			CanAccessClosed bool   `json:"can_access_closed"`
			Online          int    `json:"online"`
		} `json:"items"`
	} `json:"response"`
}

type ResponseFriends struct {
	Response int `json:"response"`
}

type GetRequests struct {
	Response struct {
		Count int `json:"count"`
		Items []struct {
			Id              int    `json:"id"`
			FirstName       string `json:"first_name"`
			LastName        string `json:"last_name"`
			IsClosed        bool   `json:"is_closed"`
			CanAccessClosed bool   `json:"can_access_closed"`
			Online          int    `json:"online"`
			UserId          int    `json:"user_id"`
		} `json:"items"`
	} `json:"response"`
}

type AreFriends struct {
	Response []struct {
		FriendStatus int `json:"friend_status"`
		UserId       int `json:"user_id"`
	} `json:"response"`
}

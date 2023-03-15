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

type ovkError struct {
	ErrorCode int    `json:"error_code"`
	ErrorMsg  string `json:"error_msg"`
	//	RequestParams []RequestParams[any] `json:"request_params"`
}

type ErrorReturned struct {
	OvkError *ovkError
	Err      error
}

//type RequestParams[T any] struct {
//	Key   string `json:"key"`
//	Value T      `json:"value"`
//}

type profileInfoResponse struct {
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
}

type ProfileInfo struct {
	profileInfoResponse `json:"response"`
}

type infoResponse struct {
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
}

type Info struct {
	infoResponse `json:"response"`
}

type countersResponse struct {
	Friends       int `json:"friends"`
	Notifications int `json:"notifications"`
	Messages      int `json:"messages"`
}

type Counters struct {
	countersResponse `json:"response"`
}

type getFriendsResponse struct {
	Count int `json:"count"`
	Items []struct {
		Id              int    `json:"id"`
		FirstName       string `json:"first_name"`
		LastName        string `json:"last_name"`
		IsClosed        bool   `json:"is_closed"`
		CanAccessClosed bool   `json:"can_access_closed"`
		Online          int    `json:"online"`
	} `json:"items"`
}

type GetFriends struct {
	getFriendsResponse `json:"response"`
}

type ResponseFriends struct {
	Response int `json:"response"`
}

type getRequestsResponse struct {
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
}

type GetRequests struct {
	getRequestsResponse `json:"response"`
}

type areFriendsResponse []struct {
	FriendStatus int `json:"friend_status"`
	UserId       int `json:"user_id"`
}

type AreFriends struct {
	areFriendsResponse `json:"response"`
}

type groupData struct {
	Id              int     `json:"id"`
	Name            string  `json:"name"`
	ScreenName      *string `json:"screen_name"`
	IsClosed        bool    `json:"is_closed"`
	CanAccessClosed bool    `json:"can_access_closed"`
	Type            string  `json:"type"`
	Verified        *int    `json:"verified"`
	HasPhoto        *int    `json:"has_photo"`
	PhotoMaxOrig    *string `json:"photo_max_orig"`
	PhotoMax        *string `json:"photo_max"`
	Photo50         *string `json:"photo_50"`
	Photo100        *string `json:"photo_100"`
	Photo200        *string `json:"photo_200"`
	Photo200Orig    *string `json:"photo_200_orig"`
	Photo400Orig    *string `json:"photo_400_orig"`
	MembersCount    *int    `json:"members_count"`
	Site            *string `json:"site"`
	Description     *string `json:"description"`
	Contacts        *[]struct {
		UserId      int    `json:"user_id"`
		Description string `json:"desc"`
	} `json:"contacts"`
	CanPost  *bool `json:"can_post"`
	IsMember *bool `json:"is_member"`
}

type getGroupsResponse struct {
	Count int         `json:"count"`
	Items []groupData `json:"items"`
}

type GetGroups struct {
	getGroupsResponse `json:"response"`
}

type GetById struct {
	Groups []groupData `json:"response"`
}

type Search struct {
	getGroupsResponse `json:"response"`
}

type JoinLeave struct {
	Status int `json:"response"`
}

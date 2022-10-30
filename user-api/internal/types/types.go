// Code generated by goctl. DO NOT EDIT.
package types

type UserInfoGetRequest struct {
	UserId int64  `form:"userId" validate:"required,gte=0" label:"用戶id"`
	Name   string `form:"name" validate:"omitempty,gte=4" label:"姓名"`
}

type UserInfoGetResponse struct {
	UserId   int64  `json:"userId"`
	Nickname string `json:"nickname"`
}

type UserInfoPostRequest struct {
	UserId int64 `json:"userId"`
}

type UserInfoPostResponse struct {
	UserId   int64  `json:"userId"`
	Nickname string `json:"nickname"`
}

type UserUpdateRequest struct {
	UserId   int64  `json:"userId"`
	Nickname string `json:"nickname"`
}

type UserUpdateResponse struct {
	Flag bool `json:"flag"`
}

type UserDeleteRequest struct {
	UserId int64 `json:"userId"`
}

type UserDeleteResponse struct {
	Flag bool `json:"flag"`
}

type RegisterReq struct {
	Mobile   string `json:"mobile"`
	Name     string `json:"name"`
	Age      int64  `json:"age"`
	Gender   int64  `json:"gender"`
	Password string `json:"password"`
}

type RegisterResp struct {
	Id           int64  `json:"id"`
	AccessToken  string `json:"accessToken"`
	AccessExpire int64  `json:"accessExpire"`
	RefreshAfter int64  `json:"refreshAfter"`
}

type UserInfoGetRpcRequest struct {
	UserId int64  `form:"userId"`
	Name   string `form:"name"`
}

type UserInfoGetRpcResponse struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

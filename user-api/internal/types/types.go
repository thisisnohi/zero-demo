// Code generated by goctl. DO NOT EDIT.
package types

type UserInfoReq struct {
	UserId int64 `json:"userId"`
}

type UserInfoResp struct {
	UserId   int64  `json:"userId"`
	NickName string `json:"string"`
}

type UserCreateReq struct {
	Mobile   string `json:"mobile"`
	Nickname string `json:"nickname"`
}

type UserCreateResp struct {
	Flag   bool `json:"flag"`
	UserId int64  `json:"userId"`
}

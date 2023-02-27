package model

type WeChatLoginReq struct {
	Code string `json:"code" validate:"min=1"`
}

type WeChatLoginResp struct {
	Token string `json:"token"`
}

type WXSessionRet struct {
	OpenId     string `json:"openid"`
	SessionKey string `json:"session_key"`
	UnionId    string `json:"unionid"`
	ErrCode    int    `json:"errcode"`
	ErrMsg     string `json:"errmsg"`
}


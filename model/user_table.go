package model

import (
	"time"
)

const (
	GenderMan   = 1
	GenderWoman = 2
)

const (
	TableNameUserProfile = "user_profile"
)

type UserProfileTable struct {
	BaseModel
	Uid        int64     // 用户ID
	Name       string    // 用户昵称
	Gender     int       // 性别
	Birth      time.Time // 生日
	Portrait   string    // 头像
	Hometown   string    // 家乡
	Openid     string    // WX用户openid
	SessionKey string    // session_key
}

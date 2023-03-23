package model

const (
	GenderMan   = 1
	GenderWoman = 2
)

const (
	TableNameUserProfile = "user_profile"
	DefaultStudyNum      = 100
)

type UserProfileTable struct {
	BaseModel
	Name             string // 用户昵称
	Gender           int    // 性别
	Portrait         string // 头像
	Openid           string // WX用户openid
	SessionKey       string // session_key
	CurrentSubjectId int64  // 当前学习题库
	StudyTotal       int    // 累计学习天数
	StudyNum         int    // 今日学习数量
}

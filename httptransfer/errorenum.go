package httptransfer

import "encoding/json"

type ErrorCode struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

func (err ErrorCode) Error() string {
	data, _ := json.Marshal(err)
	return string(data)
}

func (err ErrorCode) HasError() bool {
	return err.Code != 0
}

// common
var (
	ErrorInvalidArgument = ErrorCode{Code: 100000, Msg: "invalid params"}
)

// login
var (
	ErrorLoginForbid  = ErrorCode{Code: 101000, Msg: "账号禁止登陆，请向管理员提交反馈信息"}
	ErrorInvaildToken = ErrorCode{Code: 101001, Msg: "Token已失效，请重新登陆"}
)

// highlight exam
var ()

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
	ErrorPassportExist = ErrorCode{Code: 200, Msg: "passport exist"}
)

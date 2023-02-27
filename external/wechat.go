package external

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/li-zeyuan/common/httptransfer"
	"github.com/li-zeyuan/common/model"
	"github.com/li-zeyuan/common/mylogger"
	"go.uber.org/zap"
)

const (
	baseWXSessionUrl = "https://api.weixin.qq.com/sns/jscode2session?appid=%s&secret=%s&js_code=%s&grant_type=authorization_code"
)

type Wechat struct {
	AppId  string
	Secret string
}

func NewWechat(appId, secret string) *Wechat {
	return &Wechat{
		AppId:  appId,
		Secret: secret,
	}
}

func (w *Wechat) QueryWxSession(code string) (*model.WXSessionRet, error) {
	url := fmt.Sprintf(baseWXSessionUrl, w.AppId, w.Secret, code)
	resp, err := http.Get(url)
	if err != nil {
		mylogger.Error("get weChat session error: ", zap.Error(err))
		return nil, err
	}
	defer resp.Body.Close()

	// 解析http请求中body 数据到我们定义的结构体中
	wxResp := model.WXSessionRet{}
	decoder := json.NewDecoder(resp.Body)
	if err := decoder.Decode(&wxResp); err != nil {
		return nil, err
	}

	// 判断微信接口返回的是否是一个异常情况
	if wxResp.ErrCode != 0 {
		mylogger.Error("weChat session code error: ", zap.String("err_msg", wxResp.ErrMsg))
		return nil, httptransfer.ErrorCode{Code: wxResp.ErrCode, Msg: wxResp.ErrMsg}
	}

	return &wxResp, nil
}

package httptransfer

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/li-zeyuan/common/mylogger"
	"github.com/li-zeyuan/common/utils"
)

type uIdCtxKey string

var uId uIdCtxKey = "uid"

func JwtMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenStr := c.GetHeader("Authorization")
		token, err := jwt.ParseWithClaims(tokenStr, &utils.JwtClaims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(utils.SecretKey), nil
		})
		if err != nil {
			ErrJSONResp(c,http.StatusForbidden, err)
			return
		}

		claims, ok := token.Claims.(*utils.JwtClaims)
		if ok && token.Valid {
			mylogger.Infof("uid: %d", claims.Uid)
		} else {
			ErrJSONResp(c,http.StatusForbidden, err)
			return
		}

		ctx := c.Request.Context()
		newCtx := context.WithValue(ctx, uId, claims.Uid)
		c.Request.WithContext(newCtx)

		c.Next()
	}
}

func GetUid(c context.Context) int64 {
	if c == nil {
		mylogger.Error("content is nil")
		return 0
	}

	uid, ok := c.Value(uId).(int64)
	if !ok {
		mylogger.Error("can not transfer uIdCtxKey")
		return 0
	}

	return uid
}

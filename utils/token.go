package utils

import (
	"errors"
	"os"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
)

const secretEnvKey = "moon_jwt_secret"

var SecretKey = os.Getenv(secretEnvKey)

type JwtClaims struct {
	Uid            int64
	StandardClaims jwt.StandardClaims
}

func (j JwtClaims) Valid() error {
	if len(strconv.FormatInt(j.Uid, 10)) != 18 {
		return errors.New("uid string error")
	}

	if err := j.StandardClaims.Valid(); err != nil {
		return err
	}

	return nil
}

func GenerateToken(uid int64, expireDuration time.Duration) (string, error) {
	expire := time.Now().Add(expireDuration)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, JwtClaims{
		Uid:  uid,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expire.Unix(),
		},
	})

	return token.SignedString([]byte(SecretKey))
}

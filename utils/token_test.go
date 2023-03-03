package utils

import (
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGenerateToken(t *testing.T) {
	token, err := GenerateToken(100000000000000000, time.Hour*24*30, os.Getenv("jwt_secret"))
	assert.Equal(t, err, nil)
	t.Log(token)
}

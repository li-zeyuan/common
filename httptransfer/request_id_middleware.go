package httptransfer

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/li-zeyuan/common/utils"
)


type config struct {
	handler Handler
}

func RequestIdMiddleware(opts ...Option) gin.HandlerFunc {
	cfg := &config{}

	for _, opt := range opts {
		opt(cfg)
	}

	return func(c *gin.Context) {
		rid := c.GetHeader(utils.XRequestIDKey)
		if rid == "" {
			rid = uuid.New().String()
			c.Request.Header.Add(utils.XRequestIDKey, rid)
		}

		if cfg.handler != nil {
			cfg.handler(c, rid)
		}

		ctx := context.WithValue(c.Request.Context(), utils.XRequestIDKey, rid)
		c.Request = c.Request.WithContext(ctx)

		c.Header(utils.XRequestIDKey, rid)
		c.Next()
	}
}

type Option func(*config)

type Handler func(c *gin.Context, requestID string)

func WithHandler(handler Handler) Option {
	return func(cfg *config) {
		cfg.handler = handler
	}
}

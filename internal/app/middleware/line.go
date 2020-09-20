package middleware

import (
	"github.com/chuross/taisho/internal/app/ext/line"
	"github.com/gin-gonic/gin"
	"github.com/line/line-bot-sdk-go/linebot"
	"golang.org/x/xerrors"
)

func ValidateLineSignature(c *gin.Context) {
	client, err := line.NewClient()
	if err != nil {
		c.AbortWithError(400, xerrors.Errorf("line client instansiate failed.: %w", err))
		return
	}

	if _, err := client.ParseRequest(c.Request); err == linebot.ErrInvalidSignature {
		c.AbortWithError(400, xerrors.Errorf("invalid line signature: %w", err))
		return
	}

	c.Next()
}

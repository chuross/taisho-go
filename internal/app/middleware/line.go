package middleware

import (
	gin_ext "github.com/chuross/taisho/internal/ext/gin"
	line "github.com/chuross/taisho/internal/ext/line"
	"github.com/gin-gonic/gin"
	"github.com/line/line-bot-sdk-go/linebot"
	"golang.org/x/xerrors"
)

func ValidateLineSignature(c *gin.Context) {
	client, err := line.NewClient()
	if err != nil {
		c.AbortWithError(500, xerrors.Errorf("line client init failed: %w", err))
		return
	}

	events, err := client.ParseRequest(c.Request)
	if err == linebot.ErrInvalidSignature {
		c.AbortWithError(400, xerrors.Errorf("invalid line signature: %w", err))
		return
	}

	gin_ext.SetLineEvents(c, events)

	c.Next()
}

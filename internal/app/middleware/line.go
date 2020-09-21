package middleware

import (
	ginext "github.com/chuross/taisho/internal/app/ext/gin"
	"github.com/chuross/taisho/internal/app/ext/line"
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

	ginext.SetLineEvents(c, events)

	c.Next()
}

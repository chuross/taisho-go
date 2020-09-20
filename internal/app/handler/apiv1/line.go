package apiv1

import (
	"github.com/chuross/taisho/internal/app/ext/line"
	"github.com/gin-gonic/gin"
	"github.com/line/line-bot-sdk-go/linebot"
	"golang.org/x/xerrors"
)

func PostLineCallback(c *gin.Context) {
	client, err := line.NewClient()
	if err != nil {
		c.AbortWithError(500, xerrors.Errorf("line client init failed.: %w", err))
		return
	}
	events, err := client.ParseRequest(c.Request)
	if err != nil {
		c.AbortWithError(400, xerrors.Errorf("parse line request failed: %w", err))
		return
	}

	for _, event := range events {
		m := linebot.NewTextMessage("pong")
		client.ReplyMessage(event.ReplyToken, m)
	}

	c.Status(200)
}

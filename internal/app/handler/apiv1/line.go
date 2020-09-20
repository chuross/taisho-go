package apiv1

import (
	ginext "github.com/chuross/taisho/internal/app/ext/gin"
	"github.com/chuross/taisho/internal/app/ext/line"
	service "github.com/chuross/taisho/pkg/service/line"
	"github.com/gin-gonic/gin"
	"golang.org/x/xerrors"
)

func PostLineCallback(c *gin.Context) {
	client, err := line.NewClient()
	if err != nil {
		c.AbortWithError(500, xerrors.Errorf("line client init failed.: %w", err))
		return
	}
	events, err := ginext.GetLineEvents(c)
	if err != nil {
		c.AbortWithError(400, xerrors.Errorf("parse line request failed: %w", err))
		return
	}

	for _, event := range events {
		if ms, err := service.ReplyLineMessages(event); err != nil {
			c.AbortWithError(500, xerrors.Errorf("line message handle error: %w", err))
		} else {
			client.ReplyMessage(event.ReplyToken, ms...)
		}
	}

	c.Status(200)
}

package apiv1

import (
	gin_ext "github.com/chuross/taisho/internal/ext/gin"
	"github.com/chuross/taisho/internal/ext/line"
	service "github.com/chuross/taisho/pkg/service/line"
	"github.com/gin-gonic/gin"
	"golang.org/x/xerrors"
)

func PostLineCallback(c *gin.Context) {
	client, err := line.NewClient()
	if err != nil {
		c.AbortWithError(500, xerrors.Errorf("line client init failed: %w", err))
		return
	}
	events, err := gin_ext.GetLineEvents(c)
	if err != nil {
		c.AbortWithError(400, xerrors.Errorf("parse line request failed: %w", err))
		return
	}

	for _, event := range events {
		ms, err := service.ReplyLineMessages(c.Request.Context(), event)
		if err != nil {
			c.AbortWithError(500, xerrors.Errorf("line message handle error: %w", err))
		}
		if len(ms) > 0 {
			client.ReplyMessage(event.ReplyToken, ms...).WithContext(c.Request.Context()).Do()
		}
	}

	c.Status(200)
}

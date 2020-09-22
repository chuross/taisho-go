package gin_ext

import (
	"github.com/gin-gonic/gin"
	"github.com/line/line-bot-sdk-go/linebot"
	"golang.org/x/xerrors"
)

const (
	contextLineEventsKey = "context_line_events"
)

func GetLineEvents(c *gin.Context) ([]*linebot.Event, error) {
	events, exists := c.Get(contextLineEventsKey)
	if !exists {
		return make([]*linebot.Event, 0), xerrors.New("illegal state: line event not found")
	}
	return events.([]*linebot.Event), nil
}

func SetLineEvents(c *gin.Context, events []*linebot.Event) {
	c.Set(contextLineEventsKey, events)
}

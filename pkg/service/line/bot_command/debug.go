package bot_command

import (
	"context"
	"regexp"

	"github.com/line/line-bot-sdk-go/linebot"
)

type Debug struct {
}

func (c *Debug) Doc() string {
	return ""
}

func (c *Debug) Pattern() *regexp.Regexp {
	return regexp.MustCompile("^大将！デバッグ$")
}

func (c *Debug) Exec(ctx context.Context, event *linebot.Event, message *linebot.TextMessage) ([]linebot.SendingMessage, error) {
	return []linebot.SendingMessage{
		linebot.NewTextMessage("受け付けたよ！"),
	}, nil
}

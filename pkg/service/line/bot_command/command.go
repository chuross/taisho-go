package bot_command

import (
	"context"
	"regexp"

	"github.com/line/line-bot-sdk-go/linebot"
)

type LineBotCommand interface {
	Doc() string
	Pattern() *regexp.Regexp
	Exec(ctx context.Context, event *linebot.Event, message *linebot.TextMessage) ([]linebot.SendingMessage, error)
}

func ParseOptions(message *linebot.TextMessage) map[string]string {
	options := map[string]string{}
	return options
}

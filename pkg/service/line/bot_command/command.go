package bot_command

import (
	"context"
	"regexp"
	"strings"

	"github.com/line/line-bot-sdk-go/linebot"
)

type LineBotCommand interface {
	Doc() string
	Pattern() *regexp.Regexp
	Exec(ctx context.Context, event *linebot.Event, message *linebot.TextMessage) ([]linebot.SendingMessage, error)
}

func ParseOptions(message *linebot.TextMessage) []string {
	args := strings.Split(message.Text, " ")
	if len(args) == 1 {
		return make([]string, 0)
	}
	return args[1:]
}

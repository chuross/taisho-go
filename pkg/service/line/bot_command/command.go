package bot_command

import "github.com/line/line-bot-sdk-go/linebot"

type LineBotCommand interface {
	Doc() string
	IsExecutable(event *linebot.Event, message *linebot.TextMessage) bool
	Exec(event *linebot.Event, message *linebot.TextMessage) ([]linebot.SendingMessage, error)
}

func ParseOptions(message *linebot.TextMessage) map[string]string {
	options := map[string]string{}
	return options
}

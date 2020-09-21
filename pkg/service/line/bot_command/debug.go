package bot_command

import (
	"regexp"

	"github.com/line/line-bot-sdk-go/linebot"
)

var commandPattern = regexp.MustCompile("^大将！デバッグ$")

type Debug struct {
}

func (c *Debug) Doc() string {
	return ""
}

func (c *Debug) IsExecutable(event *linebot.Event, message *linebot.TextMessage) bool {
	return commandPattern.MatchString(message.Text)
}

func (c *Debug) Exec(event *linebot.Event, message *linebot.TextMessage) ([]linebot.SendingMessage, error) {
	return []linebot.SendingMessage{
		linebot.NewTextMessage("受け付けたよ！"),
	}, nil
}

package bot_command

import (
	"context"
	"regexp"

	"github.com/chuross/taisho/pkg/service"
	"github.com/line/line-bot-sdk-go/linebot"
	"golang.org/x/xerrors"
)

type Weather struct {
}

func (c *Weather) Doc() string {
	return "大将！アメッシュ\n関東の天気図を探してくるよ"
}

func (c *Weather) Pattern() *regexp.Regexp {
	return regexp.MustCompile("^大将！アメッシュ$")
}

func (c *Weather) Exec(ctx context.Context, event *linebot.Event, message *linebot.TextMessage) ([]linebot.SendingMessage, error) {
	imageUrl, err := service.GetWeatherImage()
	if err != nil {
		return make([]linebot.SendingMessage, 0), xerrors.Errorf("weather command failed: %w", err)
	}

	return []linebot.SendingMessage{
		linebot.NewTextMessage("へいお待ち！"),
		linebot.NewImageMessage(*imageUrl, *imageUrl),
	}, nil
}

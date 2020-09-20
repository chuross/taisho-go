package line

import (
	"github.com/chuross/taisho/pkg/config"
	"github.com/line/line-bot-sdk-go/linebot"
)

func NewClient() (*linebot.Client, error) {
	c := config.Get()
	return linebot.New(c.Line.ClientSecret, c.Line.ClientAccessToken)
}

package bot_command

import (
	"context"
	"regexp"

	"github.com/chuross/taisho/pkg/service"
	"github.com/line/line-bot-sdk-go/linebot"
	"golang.org/x/xerrors"
)

type Itsumono struct {
}

func (c *Itsumono) Doc() string {
	return "大将！いつもの - いつもの握ってまってやすぜ"
}

func (c *Itsumono) Pattern() *regexp.Regexp {
	return regexp.MustCompile("^大将！いつもの$")
}

func (c *Itsumono) Exec(ctx context.Context, event *linebot.Event, message *linebot.TextMessage) ([]linebot.SendingMessage, error) {
	imageUrl, err := service.Itsumono(ctx)
	if err != nil || imageUrl == nil {
		return make([]linebot.SendingMessage, 0), xerrors.Errorf("itsumono command failed: %w", err)
	}

	return []linebot.SendingMessage{
		linebot.NewTextMessage("へいお待ち！"),
		linebot.NewImageMessage(*imageUrl, *imageUrl),
	}, nil
}

package bot_command

import (
	"context"
	"log"
	"regexp"

	"github.com/line/line-bot-sdk-go/linebot"
)

type Debug struct {
}

func (c *Debug) Doc() string {
	return ""
}

func (c *Debug) Pattern() *regexp.Regexp {
	return regexp.MustCompile(`^大将！デバッグ$`)
}

func (c *Debug) Exec(ctx context.Context, event *linebot.Event, message *linebot.TextMessage) ([]linebot.SendingMessage, error) {
	log.Printf("line_user_id=%s, group_id=%s, room_id=%s", event.Source.UserID, event.Source.GroupID, event.Source.RoomID)

	return []linebot.SendingMessage{
		linebot.NewTextMessage("受け付けたよ！"),
	}, nil
}

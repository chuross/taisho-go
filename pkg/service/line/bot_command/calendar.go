package bot_command

import (
	"context"
	"os"
	"regexp"
	"strings"

	"github.com/chuross/taisho/pkg/service"
	"github.com/chuross/taisho/pkg/util"
	"github.com/line/line-bot-sdk-go/linebot"
	"golang.org/x/xerrors"
	"google.golang.org/api/calendar/v3"
)

type Calendar struct {
}

func (c *Calendar) Doc() string {
	return "大将！予定 <yyyy-MM-dd> <内容>\nGoogleカレンダーに登録するよ！(管理者のみ)"
}

func (c *Calendar) Pattern() *regexp.Regexp {
	return regexp.MustCompile("^大将！予定 ([0-9]{4}-[0-9]{2}-[0-9]{2}) (.+)")
}

func (c *Calendar) Exec(ctx context.Context, event *linebot.Event, message *linebot.TextMessage) ([]linebot.SendingMessage, error) {
	wlgID := os.Getenv("TAISHO_WHITE_LIST_GROUP_ID")
	wlgIDs := strings.Split(wlgID, ",")

	if !util.ContainsString(wlgIDs, event.Source.GroupID) {
		return []linebot.SendingMessage{
			linebot.NewTextMessage("この注文は一見さんはお断りだよ"),
		}, nil
	}

	options := ParseOptions(message)
	if len(options) != 2 {
		return []linebot.SendingMessage{}, xerrors.New("invalid arguments")
	}

	date := options[0]
	body := options[1]

	e := &calendar.Event{
		Summary: body,
		Start: &calendar.EventDateTime{
			Date: date,
		},
	}

	if err := service.CreateCalendarEvent(ctx, os.Getenv("TAISHO_CALENDAR_ID"), e); err != nil {
		return []linebot.SendingMessage{}, xerrors.Errorf("create calendar event failed: %w", err)
	}

	return []linebot.SendingMessage{
		linebot.NewTextMessage("カレンダーに登録したよ！確認してくんな！"),
		linebot.NewTextMessage(date + " " + body),
	}, nil
}

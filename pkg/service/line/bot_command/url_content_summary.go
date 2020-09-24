package bot_command

import (
	"context"
	"log"
	"regexp"
	"strings"

	"github.com/chuross/taisho/pkg/service"
	"github.com/line/line-bot-sdk-go/linebot"
	"golang.org/x/xerrors"
)

var urlContentSummaryPattern = regexp.MustCompile(`https?://[\w!\?/\+\-_~=;\.,\*&@#\$%\(\)'\[\]]+`)

type UrlContentSummary struct {
}

func (c *UrlContentSummary) Doc() string {
	return ""
}

func (c *UrlContentSummary) Pattern() *regexp.Regexp {
	return urlContentSummaryPattern
}

func (c *UrlContentSummary) Exec(ctx context.Context, event *linebot.Event, message *linebot.TextMessage) ([]linebot.SendingMessage, error) {
	url := urlContentSummaryPattern.FindString(message.Text)
	summary, err := service.GetUrlContentSummary(url)
	if err != nil {
		return make([]linebot.SendingMessage, 0), xerrors.Errorf("url content summary command failed: %w", err)
	}

	if summary == nil {
		return make([]linebot.SendingMessage, 0), nil
	}

	if len(summary.Summaries) == 0 {
		log.Printf("empty summary url=" + url)
		return make([]linebot.SendingMessage, 0), nil
	}

	var text strings.Builder
	text.WriteString("要約しといたよ！" + "\n")
	text.WriteString(url + "\n")
	text.WriteString(strings.Join(summary.Summaries, "\n\n"))

	return []linebot.SendingMessage{
		linebot.NewTextMessage(text.String()),
	}, nil
}

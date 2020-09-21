package service

import (
	"strings"

	"github.com/chuross/taisho/pkg/service/line/bot_command"
	"github.com/line/line-bot-sdk-go/linebot"
	"golang.org/x/xerrors"
)

var (
	commands = []bot_command.LineBotCommand{
		&bot_command.Debug{},
	}
)

func ReplyLineMessages(event *linebot.Event) ([]linebot.SendingMessage, error) {
	messages := make([]linebot.SendingMessage, 0)
	switch event.Type {
	case linebot.EventTypeMessage:
		switch m := event.Message.(type) {
		case *linebot.TextMessage:
			return dispatch(event, m)
		}
	}
	return messages, nil
}

func dispatch(event *linebot.Event, message *linebot.TextMessage) ([]linebot.SendingMessage, error) {
	if message.Text == "大将！" {
		return helpCommand(), nil
	}

	messages := make([]linebot.SendingMessage, 0)
	for _, command := range commands {
		if !command.IsExecutable(event, message) {
			continue
		}
		ms, err := command.Exec(event, message)
		if err != nil {
			return messages, xerrors.Errorf("line reply error: %w", err)
		}
		messages = append(messages, ms...)
	}

	return messages, nil
}

func helpCommand() []linebot.SendingMessage {
	docs := make([]string, 0)
	docs = append(docs, "以下の注文を受け付けているよ！")
	for _, c := range commands {
		if len(c.Doc()) > 0 {
			docs = append(docs, c.Doc())
		}
	}

	return []linebot.SendingMessage{
		linebot.NewTextMessage(strings.Join(docs, "\n\n")),
	}
}

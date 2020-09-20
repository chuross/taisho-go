package service

import "github.com/line/line-bot-sdk-go/linebot"

func ReplyLineMessages() ([]linebot.SendingMessage, error) {
	messages := make([]linebot.SendingMessage, 0)
	return messages, nil
}

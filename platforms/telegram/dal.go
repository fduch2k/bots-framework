package telegram_bot

import (
	"golang.org/x/net/context"
	"github.com/strongo/db"
)

type TgChatInstanceDal interface {
	GetTelegramChatInstanceByID(c context.Context, id string) (tgChatInstance TelegramChatInstance, err error)
	NewTelegramChatInstance(chatInstanceID string, chatID int64, preferredLanguage string) (tgChatInstance TelegramChatInstance)
	SaveTelegramChatInstance(c context.Context, tgChatInstance TelegramChatInstance) (err error)
}

type dal struct {
	DB db.Database
	TgChatInstance TgChatInstanceDal
}

var DAL dal

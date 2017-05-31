package telegram_bot

import (
	"fmt"
	"github.com/strongo/bots-framework/core"
	"time"
	"strconv"
	"google.golang.org/appengine/datastore"
	"github.com/strongo/app/gaedb"
)

const (
	TelegramChatKind = "TgChat"
)


type TelegramChat struct {
	bots.BotChatEntity
	TelegramUserID        int
	LastProcessedUpdateID int `datastore:",noindex"`
}

var _ bots.BotChat = (*TelegramChat)(nil)

func NewTelegramChat() TelegramChat {
	return TelegramChat{
		BotChatEntity: bots.BotChatEntity{
			BotEntity: bots.BotEntity{
				OwnedByUser: bots.OwnedByUser{
					DtCreated: time.Now(),
				},
			},
		},
	}
}

func (entity *TelegramChat) SetAppUserIntID(id int64) {
	entity.AppUserIntID = id
}

func (entity *TelegramChat) SetBotUserID(id interface{}) {
	switch id.(type) {
	case string:
		var err error
		entity.TelegramUserID, err = strconv.Atoi(id.(string))
		if err != nil {
			panic(err.Error())
		}
	case int:
		entity.TelegramUserID = id.(int)
	case int64:
		entity.TelegramUserID = id.(int)
	default:
		panic(fmt.Sprintf("Expected string, got: %T=%v", id, id))
	}
}

func (entity *TelegramChat) Load(ps []datastore.Property) error {
	return datastore.LoadStruct(entity, ps)
}

func (entity *TelegramChat) Save() (properties []datastore.Property, err error) {
	if properties, err = datastore.SaveStruct(entity); err != nil {
		return properties, err
	}

	if properties, err = gaedb.CleanProperties(properties, map[string]gaedb.IsOkToRemove{
		"AccessGranted":         gaedb.IsFalse,
		"AwaitingReplyTo":       gaedb.IsEmptyString,
		"DtForbidden":           gaedb.IsZeroTime,
		"DtForbiddenLast":       gaedb.IsZeroTime,
		"GaClientID":            gaedb.IsEmptyByteArray,
		"LastProcessedUpdateID": gaedb.IsZeroInt,
		"PreferredLanguage":     gaedb.IsEmptyString,
		"Title":                 gaedb.IsEmptyString,  // TODO: Is it obsolete?
		"Type":                  gaedb.IsEmptyString,  // TODO: Is it obsolete?
	}); err != nil {
		return
	}

	return
}

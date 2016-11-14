package viber_bot

import (
	"github.com/strongo/bots-framework/core"
)

type ViberWebhookChat struct {
	viberUserID string
}

var _ bots.WebhookChat = (*ViberWebhookChat)(nil)

func (wh ViberWebhookChat) GetID() interface{} {
	return wh.viberUserID
}

func (wh ViberWebhookChat) GetFullName() string {
	return ""
}

func (wh ViberWebhookChat) GetType() string {
	return "private"
}

func NewViberWebhookChat(viberUserID string) ViberWebhookChat {
	return ViberWebhookChat{viberUserID: viberUserID}
}
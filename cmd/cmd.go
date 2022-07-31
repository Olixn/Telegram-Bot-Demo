package cmd

import (
	"github.com/Olixn/Telegram-Bot-Demo/plugins/yiyan"
	"github.com/Olixn/Telegram-Bot-Demo/utils"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Bot struct {
	Bot *tgbotapi.BotAPI
	Ctx tgbotapi.Update
}

func New(bot *tgbotapi.BotAPI, ctx tgbotapi.Update) *Bot {
	return &Bot{
		Bot: bot,
		Ctx: ctx,
	}
}

func (b *Bot) Handler() {
	msg := b.Ctx.Message.Text
	cmd := utils.RegexpCMD(msg, " ")
	if cmd == "" {
		cmd = msg
	}
	switch cmd {
	case "一言":
		reply, _ := yiyan.GetYiYan()
		b.SendMsg(reply)
		break
	}
}

func (b *Bot) SendMsg(reply string) {
	msg := tgbotapi.NewMessage(b.Ctx.Message.Chat.ID, reply)
	msg.ReplyToMessageID = b.Ctx.Message.MessageID

	b.Bot.Send(msg)
}

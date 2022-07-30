package main

import (
	"fmt"
	"github.com/Olixn/Telegram-Bot-Demo/config"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"net/http"
	"net/url"
)

var (
	AppConfig   *config.Config
	ApiEndPoint = "https://api.telegram.org/bot%s/%s"
)

func init() {
	AppConfig = config.InitConfig()
	fmt.Println(AppConfig)
}

func main() {
	var bot *tgbotapi.BotAPI
	var err error
	fmt.Println(AppConfig.Bot.Token, AppConfig.Bot.EnableProxy)
	if AppConfig.Bot.EnableProxy {
		proxy, _ := url.Parse(AppConfig.Bot.Proxy)
		bot, err = tgbotapi.NewBotAPIWithClient(AppConfig.Bot.Token, ApiEndPoint, &http.Client{
			Transport: &http.Transport{
				Proxy: http.ProxyURL(proxy),
			},
		})
	} else {
		bot, err = tgbotapi.NewBotAPI(AppConfig.Bot.Token)
	}

	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil {
			log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

			msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
			msg.ReplyToMessageID = update.Message.MessageID

			bot.Send(msg)
		}
	}
}

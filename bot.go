package telegram

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

func SendTGMsg(msg, token string, groupId int64) {
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	//updates, err := bot.GetUpdatesChan(u)

	bot.Send(tgbotapi.NewMessage(groupId, msg))

}

package telegram

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
)

func GetBot(token string) (*tgbotapi.BotAPI, error) {
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		return nil, err
	}

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	return bot, nil

}

func SendMsg(bot *tgbotapi.BotAPI, msg string, groupId int64) (tgbotapi.Message, error) {
	return bot.Send(tgbotapi.NewMessage(groupId, msg))
}

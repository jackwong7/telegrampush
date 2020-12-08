package telegram

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
)

func GetBot(token string) (*tgbotapi.BotAPI, error) {
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		return nil, err
	}

	return bot, nil

}

func SendMsg(bot *tgbotapi.BotAPI, msg string, groupId int64) (tgbotapi.Message, error) {
	return bot.Send(tgbotapi.NewMessage(groupId, msg))
}

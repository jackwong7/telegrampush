package telegram

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
)

//获取bot实例
func GetBot(token string) (*tgbotapi.BotAPI, error) {
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		return nil, err
	}
	return bot, nil
}

//发送消息
func SendMsg(bot *tgbotapi.BotAPI, msg string, groupId int64) (tgbotapi.Message, error) {
	return bot.Send(tgbotapi.NewMessage(groupId, msg))
}

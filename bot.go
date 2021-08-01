package telegram

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type BotConfig struct {
	Host   string
	Token  string
	UserId int64
}

//发送消息
func SendMsg(bot *BotConfig, msg string) *http.Response {
	if bot.Host == "" {
		bot.Host = "api.telegram.org"
	}
	client := &http.Client{}
	req, _ := http.NewRequest("POST", fmt.Sprintf("https://%s/bot%s/sendMessage",
		bot.Host, bot.Token), getReader(bot, msg))
	tgReq := tgReq{req: req}
	tgReq.initJsonReq()
	resp, err := client.Do(tgReq.req)
	if err != nil {
		panic(err)
	}
	return resp

}

//tg请求头
type tgReq struct {
	req *http.Request
}

//获取入参
func getReader(bot *BotConfig, msg string) *bytes.Reader {
	data := make(map[string]interface{})
	data["chat_id"] = bot.UserId
	data["text"] = msg
	bytesData, err := json.Marshal(data)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	return bytes.NewReader(bytesData)
}

//初始化tg请求头
func (req *tgReq) initJsonReq() {
	req.req.Header.Set("Content-type", "application/json")
	req.req.Header.Set("Accept", "application/json")
}

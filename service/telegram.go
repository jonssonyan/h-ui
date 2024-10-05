package service

import (
	"errors"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/sirupsen/logrus"
	"h-ui/dao"
	"h-ui/model/constant"
)

var bot *tgbotapi.BotAPI

func NewBotApi() (*tgbotapi.BotAPI, error) {
	// 参数校验
	enable, err := dao.GetConfig("key = ?", constant.TelegramEnable)
	if err != nil {
		return nil, err
	}
	if enable.Value == nil || *enable.Value != "1" {
		return nil, errors.New("telegram not enable")
	}
	token, err := dao.GetConfig("key = ?", constant.TelegramToken)
	if err != nil {
		return nil, err
	}
	if token.Value == nil || *token.Value == "" {
		return nil, errors.New("telegram token not set")
	}
	chatIdStr, err := dao.GetConfig("key = ?", constant.TelegramChatId)
	if err != nil {
		return nil, err
	}
	if chatIdStr.Value == nil || *chatIdStr.Value == "" {
		return nil, errors.New("telegram chatId not set")
	}
	// 新建对象
	if bot != nil {
		return bot, nil
	}
	bot, err = tgbotapi.NewBotAPI(*token.Value)
	if err != nil {
		logrus.Errorf("new bot api err: %v", err)
		return nil, err
	}
	bot.Debug = false
	logrus.Infof("Authorized on account %s", bot.Self.UserName)
	// 初始化菜单

	return bot, nil
}

func GetMe() (tgbotapi.User, error) {
	botApi, err := NewBotApi()
	if err != nil {
		return tgbotapi.User{}, err
	}
	user, err := botApi.GetMe()
	if err != nil {
		logrus.Errorf("tg api GetMe err: %v", err)
		return tgbotapi.User{}, err
	}
	return user, nil
}

func GetUpdatesChan() (tgbotapi.UpdatesChannel, error) {
	botApi, err := NewBotApi()
	if err != nil {
		return nil, err
	}
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	return botApi.GetUpdatesChan(u), nil
}

func SendWithMessage(chatId int64, text string) error {
	botApi, err := NewBotApi()
	if err != nil {
		return err
	}
	message := tgbotapi.NewMessage(chatId, text)
	if _, err = botApi.Send(message); err != nil {
		logrus.Errorf("tg api SendMessage err: %v chatId: %d text: %s", err, chatId, text)
		return err
	}
	return nil
}

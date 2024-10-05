package service

import (
	"errors"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/sirupsen/logrus"
	"h-ui/dao"
	"h-ui/model/constant"
	"strconv"
)

var bot *tgbotapi.BotAPI

// 参数校验
func valid() (string, int64, error) {
	enable, err := dao.GetConfig("key = ?", constant.TelegramEnable)
	if err != nil {
		return "", 0, err
	}
	if enable.Value == nil || *enable.Value != "1" {
		return "", 0, errors.New("telegram not enable")
	}
	token, err := dao.GetConfig("key = ?", constant.TelegramToken)
	if err != nil {
		return "", 0, err
	}
	if token.Value == nil || *token.Value == "" {
		return "", 0, errors.New("telegram token not set")
	}
	chatIdStr, err := dao.GetConfig("key = ?", constant.TelegramChatId)
	if err != nil {
		return "", 0, err
	}
	if chatIdStr.Value == nil || *chatIdStr.Value == "" {
		return "", 0, errors.New("telegram chatId not set")
	}
	chatId, err := strconv.ParseInt(*chatIdStr.Value, 10, 64)
	if err != nil {
		return "", 0, err
	}
	return *token.Value, chatId, nil
}

func init() {
	token, _, err := valid()
	if err != nil {
		return
	}
	bot, err = tgbotapi.NewBotAPI(token)
	if err != nil {
		logrus.Errorf("new bot api err: %v", err)
		return
	}
	bot.Debug = false
	logrus.Infof("Authorized on account %s", bot.Self.UserName)
	// 初始化 menu
	commands := []tgbotapi.BotCommand{
		{Command: "status", Description: "状态"},
	}
	setCommands := tgbotapi.NewSetMyCommands(commands...)
	if _, err := bot.Request(setCommands); err != nil {
		logrus.Errorf("unable to set commands err: %v", err)
		return
	}
}

func GetMe() (tgbotapi.User, error) {
	user, err := bot.GetMe()
	if err != nil {
		logrus.Errorf("tg api GetMe err: %v", err)
		return tgbotapi.User{}, err
	}
	return user, nil
}

func GetUpdatesChan() tgbotapi.UpdatesChannel {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	return bot.GetUpdatesChan(u)
}

func SendWithMessage(chatId int64, text string) error {
	message := tgbotapi.NewMessage(chatId, text)
	if _, err := bot.Send(message); err != nil {
		logrus.Errorf("tg api SendMessage err: %v chatId: %d text: %s", err, chatId, text)
		return err
	}
	return nil
}

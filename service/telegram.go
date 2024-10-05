package service

import (
	"errors"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/sirupsen/logrus"
	"h-ui/dao"
	"h-ui/model/constant"
)

var bot *tgbotapi.BotAPI

// 参数校验
func valid() (string, string, error) {
	enable, err := dao.GetConfig("key = ?", constant.TelegramEnable)
	if err != nil {
		return "", "", err
	}
	if enable.Value == nil || *enable.Value != "1" {
		return "", "", errors.New("telegram not enable")
	}
	token, err := dao.GetConfig("key = ?", constant.TelegramToken)
	if err != nil {
		return "", "", err
	}
	if token.Value == nil || *token.Value == "" {
		return "", "", errors.New("telegram token not set")
	}
	username, err := dao.GetConfig("key = ?", constant.TelegramUsername)
	if err != nil {
		return "", "", err
	}
	if username.Value == nil || *username.Value == "" {
		return "", "", errors.New("telegram username not set")
	}
	return *token.Value, *username.Value, nil
}

func InitTelegramBot() error {
	token, username, err := valid()
	if err != nil {
		return err
	}
	bot, err = tgbotapi.NewBotAPI(token)
	if err != nil {
		logrus.Errorf("new bot api err: %v", err)
		return err
	}
	bot.Debug = true
	logrus.Infof("Authorized on account %s", bot.Self.UserName)
	// 初始化 menu
	commands := []tgbotapi.BotCommand{
		{Command: "status", Description: "Status"},
	}
	setCommands := tgbotapi.NewSetMyCommands(commands...)
	if _, err := bot.Request(setCommands); err != nil {
		logrus.Errorf("unable to set commands err: %v", err)
		return err
	}
	// 处理消息
	go func() {
		updates := getUpdatesChan()
		for update := range updates {
			handleMsg(update, username)
		}
	}()
	return nil
}

func getUpdatesChan() tgbotapi.UpdatesChannel {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	return bot.GetUpdatesChan(u)
}

func handleMsg(update tgbotapi.Update, username string) {
	if update.Message != nil && update.Message.IsCommand() && update.SentFrom().UserName == username {
		switch update.Message.Command() {
		case "status":
			if err := handleStatus(update); err != nil {
				logrus.Errorf("handleStatus err: %v", err)
			}
		default:
			if err := handleDefault(update); err != nil {
				logrus.Errorf("handleDefault err: %v", err)
			}
		}
	}
}

func handleStatus(update tgbotapi.Update) error {
	if err := SendWithMessage(update.Message.Chat.ID, "你好"); err != nil {
		return err
	}
	return nil
}

func handleDefault(update tgbotapi.Update) error {
	return nil
}

func GetMe() (tgbotapi.User, error) {
	user, err := bot.GetMe()
	if err != nil {
		logrus.Errorf("tg api GetMe err: %v", err)
		return tgbotapi.User{}, err
	}
	return user, nil
}

func SendWithMessage(chatId int64, text string) error {
	message := tgbotapi.NewMessage(chatId, text)
	if _, err := bot.Send(message); err != nil {
		logrus.Errorf("tg api SendMessage err: %v chatId: %d text: %s", err, chatId, text)
		return err
	}
	return nil
}

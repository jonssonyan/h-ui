package service

import (
	"errors"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/sirupsen/logrus"
	"h-ui/dao"
	"h-ui/model/constant"
	"os"
	"strconv"
	"strings"
	"time"
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
	chatId, err := dao.GetConfig("key = ?", constant.TelegramChatId)
	if err != nil {
		return "", "", err
	}
	if chatId.Value == nil || *chatId.Value == "" {
		return "", "", errors.New("telegram chatId not set")
	}
	return *token.Value, *chatId.Value, nil
}

func InitTelegramBot() error {
	token, chatId, err := valid()
	if err != nil {
		return err
	}
	bot, err = tgbotapi.NewBotAPI(token)
	if err != nil {
		logrus.Errorf("new bot api err: %v", err)
		return err
	}
	bot.Debug = os.Getenv(constant.TelegramDebug) == "true"
	logrus.Infof("Authorized on account %s", bot.Self.UserName)
	// 初始化 menu
	commands := []tgbotapi.BotCommand{
		{Command: "chatId", Description: "获取chatId"},
		{Command: "status", Description: "系统状态"},
		{Command: "restart", Description: "重启系统"},
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
			handleMsg(update, chatId)
		}
	}()
	return nil
}

func getUpdatesChan() tgbotapi.UpdatesChannel {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	return bot.GetUpdatesChan(u)
}

func handleMsg(update tgbotapi.Update, chatId string) {
	if update.Message != nil && update.Message.IsCommand() && strconv.FormatInt(update.Message.Chat.ID, 10) == chatId {
		switch update.Message.Command() {
		case "chatId":
			if err := handleChatId(update); err != nil {
				logrus.Errorf("handleStatus err: %v", err)
			}
		case "status":
			if err := handleStatus(update); err != nil {
				logrus.Errorf("handleStatus err: %v", err)
			}
		case "restart":
			if err := handleRestart(update); err != nil {
				logrus.Errorf("handleRestart err: %v", err)
			}
		default:
			if err := handleDefault(update); err != nil {
				logrus.Errorf("handleDefault err: %v", err)
			}
		}
	}
}

func handleChatId(update tgbotapi.Update) error {
	if err := SendWithMessage(update.Message.Chat.ID, fmt.Sprintf("chatId: %d", update.Message.Chat.ID)); err != nil {
		return err
	}
	return nil
}

func handleStatus(update tgbotapi.Update) error {
	text := "【系统状态】\n"
	systemMonitorVo, err := MonitorSystem()
	if err != nil {
		return err
	}
	hysteria2MonitorVo, err := MonitorHysteria2()
	if err != nil {
		return err
	}
	text += fmt.Sprintf("H UI 版本: %s\n", systemMonitorVo.HUIVersion)
	text += fmt.Sprintf("CPU 使用率: %.1f%%\n", systemMonitorVo.CpuPercent)
	text += fmt.Sprintf("内存使用率: %.1f%%\n", systemMonitorVo.MemPercent)
	text += fmt.Sprintf("磁盘使用率: %.1f%%\n", systemMonitorVo.DiskPercent)

	text += fmt.Sprintf("Hysteria2 版本: %s\n", hysteria2MonitorVo.Version)
	var status = "运行"
	if !hysteria2MonitorVo.Running {
		status = "停止"
	}
	text += fmt.Sprintf("Hysteria2 状态: %s\n", status)
	text += fmt.Sprintf("在线用户数: %d\n", hysteria2MonitorVo.UserTotal)
	text += fmt.Sprintf("在线设备数: %d\n", hysteria2MonitorVo.DeviceTotal)

	if err := SendWithMessage(update.Message.Chat.ID, text); err != nil {
		return err
	}
	return nil
}

func handleRestart(update tgbotapi.Update) error {
	_ = StopServer()
	if err := SendWithMessage(update.Message.Chat.ID, "重启成功"); err != nil {
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

// TelegramLoginRemind 登录提醒
func TelegramLoginRemind(username string, ip string) {
	configs, err := dao.ListConfig("key in ?", []string{
		constant.TelegramEnable,
		constant.TelegramChatId,
		constant.TelegramLoginJobEnable,
		constant.TelegramLoginJobText})
	if err != nil {
		return
	}
	var telegramEnable, telegramChatId, telegramLoginJobEnable, telegramLoginJobText = "0", "", "0", ""
	for _, item := range configs {
		if item.Value != nil {
			key := *item.Key
			value := *item.Value
			if key == constant.TelegramEnable {
				telegramEnable = value
			} else if key == constant.TelegramChatId {
				telegramChatId = value
			} else if key == constant.TelegramLoginJobEnable {
				telegramLoginJobEnable = value
			} else if key == constant.TelegramLoginJobText {
				telegramLoginJobText = value
			}
		}
	}

	if telegramEnable != "1" || telegramChatId == "" || telegramLoginJobEnable != "1" || telegramLoginJobText == "" {
		return
	}

	chatId, err := strconv.ParseInt(telegramChatId, 10, 64)
	if err != nil {
		logrus.Errorf("parse chatId err: %v", err)
		return
	}

	telegramLoginJobText = strings.ReplaceAll(telegramLoginJobText, "[time]", time.Now().Format("20060102150405"))
	telegramLoginJobText = strings.ReplaceAll(telegramLoginJobText, "[username]", username)
	telegramLoginJobText = strings.ReplaceAll(telegramLoginJobText, "[ip]", ip)

	if err = SendWithMessage(chatId, telegramLoginJobText); err != nil {
		return
	}
}

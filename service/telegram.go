package service

import (
	"context"
	"errors"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/sirupsen/logrus"
	"h-ui/dao"
	"h-ui/model/constant"
	"net"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

var bot *tgbotapi.BotAPI

var done = make(chan bool)

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
	if chatId.Value == nil {
		return "", "", errors.New("telegram chatId is nil")
	}
	return *token.Value, *chatId.Value, nil
}

func InitTelegramBot() error {
	token, chatId, err := valid()
	if err != nil {
		if err.Error() == "telegram not enable" {
			return nil
		}
		return err
	}
	dnsServers := os.Getenv(constant.TelegramDNSServers)
	if dnsServers != "" {
		for _, s := range strings.Split(dnsServers, ",") {
			addr := strings.TrimSpace(s)
			if addr == "" {
				continue
			}
			r := &net.Resolver{PreferGo: true, Dial: func(ctx context.Context, network, address string) (net.Conn, error) {
				d := net.Dialer{}
				return d.DialContext(ctx, "udp", net.JoinHostPort(addr, "53"))
			}}
			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			_, err := r.LookupHost(ctx, "api.telegram.org")
			cancel()
			if err == nil {
				net.DefaultResolver = r
				break
			}
		}
	}
	httpClient := &http.Client{Timeout: 5 * time.Second}
	bot, err = tgbotapi.NewBotAPIWithClient(token, tgbotapi.APIEndpoint, httpClient)
	if err != nil {
		logrus.Errorf("new bot api err: %v", err)
		return err
	}
	bot.Debug = os.Getenv(constant.TelegramDebug) == "true"
	commands := []tgbotapi.BotCommand{
		{Command: "status", Description: "System Status"},
		{Command: "restart", Description: "System Restart"},
	}
	if chatId == "" {
		commands = append(commands, tgbotapi.BotCommand{Command: "chatid", Description: "Get chatId"})
	}
	setCommands := tgbotapi.NewSetMyCommands(commands...)
	if _, err := bot.Request(setCommands); err != nil {
		logrus.Errorf("unable to set commands err: %v", err)
		return err
	}
	go func(done chan bool) {
		updates := getUpdatesChan()
		for {
			select {
			case update, ok := <-updates:
				if !ok {
					return
				}
				handleMsg(update, chatId)
			case <-done:
				break
			}
		}
	}(done)
	return nil
}

func getUpdatesChan() tgbotapi.UpdatesChannel {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	return bot.GetUpdatesChan(u)
}

func handleMsg(update tgbotapi.Update, chatId string) {
	if update.Message != nil && update.Message.IsCommand() && (chatId == "" || strconv.FormatInt(update.Message.Chat.ID, 10) == chatId) {
		switch update.Message.Command() {
		case "chatid":
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
	systemMonitorVo, err := MonitorSystem()
	if err != nil {
		return err
	}
	hysteria2MonitorVo, err := MonitorHysteria2()
	if err != nil {
		return err
	}
	text := "【H UI】\n"
	text += fmt.Sprintf("H UI Version: %s\n", systemMonitorVo.HUIVersion)
	text += fmt.Sprintf("CPU Usage: %.1f%%\n", systemMonitorVo.CpuPercent)
	text += fmt.Sprintf("Memory Usage: %.1f%%\n", systemMonitorVo.MemPercent)
	text += fmt.Sprintf("Disk Usage: %.1f%%\n", systemMonitorVo.DiskPercent)

	text += fmt.Sprintf("Hysteria2 Version: %s\n", hysteria2MonitorVo.Version)
	var status = "Running"
	if !hysteria2MonitorVo.Running {
		status = "Stop"
	}
	text += fmt.Sprintf("Hysteria2 Status: %s\n", status)
	text += fmt.Sprintf("Number of online users: %d\n", hysteria2MonitorVo.UserTotal)
	text += fmt.Sprintf("Number of online devices: %d\n", hysteria2MonitorVo.DeviceTotal)

	if err := SendWithMessage(update.Message.Chat.ID, text); err != nil {
		return err
	}
	return nil
}

func handleRestart(update tgbotapi.Update) error {
	_ = StopServer()
	if err := SendWithMessage(update.Message.Chat.ID, "Restart successful"); err != nil {
		return err
	}
	done <- true
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
func TelegramLoginRemind(username string, ip string) error {
	configs, err := dao.ListConfig("key in ?", []string{
		constant.TelegramEnable,
		constant.TelegramChatId,
		constant.TelegramLoginJobEnable,
		constant.TelegramLoginJobText})
	if err != nil {
		return err
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
		return nil
	}

	chatId, err := strconv.ParseInt(telegramChatId, 10, 64)
	if err != nil {
		logrus.Errorf("parse chatId err: %v", err)
		return err
	}

	telegramLoginJobText = strings.ReplaceAll(telegramLoginJobText, "[time]", time.Now().Format("2006-01-02 15:04:05"))
	telegramLoginJobText = strings.ReplaceAll(telegramLoginJobText, "[username]", username)
	telegramLoginJobText = strings.ReplaceAll(telegramLoginJobText, "[ip]", ip)

	if err = SendWithMessage(chatId, fmt.Sprintf("【H UI】\n%s", telegramLoginJobText)); err != nil {
		return err
	}
	return nil
}

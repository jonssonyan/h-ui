package proxy

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	"h-ui/model/bo"
	"h-ui/model/constant"
	"io"
	"net/http"
	"time"
)

type Hysteria2Api struct {
	apiPort int64
}

func NewHysteria2Api(apiPort int64) *Hysteria2Api {
	return &Hysteria2Api{
		apiPort: apiPort,
	}
}

// ListUsers 每个用户的流量信息
func (n *Hysteria2Api) ListUsers(clear bool) (map[string]bo.Hysteria2UserTraffic, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	url := fmt.Sprintf("http://127.0.0.1:%d/traffic", n.apiPort)
	if clear {
		url = fmt.Sprintf("%s?clear=1", url)
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		logrus.Errorf("Hysteria2 ListUsers NewRequest err: %v", err)
		return nil, errors.New(constant.SysError)
	}
	resp, err := http.DefaultClient.Do(req)
	defer func() {
		if resp != nil {
			_ = resp.Body.Close()
		}
	}()
	if err != nil || resp.StatusCode != http.StatusOK {
		logrus.Errorf("Hysteria2 ListUsers err: %v", err)
		return nil, errors.New("http connection error")
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		logrus.Errorf("Hysteria2 io read err: %v", err)
		return nil, errors.New("http connection error")
	}
	var users map[string]bo.Hysteria2UserTraffic
	if err = json.Unmarshal(body, &users); err != nil {
		logrus.Errorf("Hysteria2 ListUsers Unmarshal err: %v", err)
		return nil, errors.New(constant.SysError)
	}
	return users, nil
}

func (n *Hysteria2Api) GetUser(conPass string, clear bool) (*bo.Hysteria2User, error) {
	users, err := n.ListUsers(clear)
	if err != nil {
		return nil, err
	}
	user := users[conPass]
	return &bo.Hysteria2User{
		ConPass: conPass,
		Tx:      user.Tx,
		Rx:      user.Rx,
	}, nil
}

// KickUsers 踢下线
func (n *Hysteria2Api) KickUsers(keys []string) error {
	usernamesByte, err := json.Marshal(keys)
	if err != nil {
		logrus.Errorf("Hysteria2 KickUsers Marshal err: %v", err)
		return errors.New(constant.SysError)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	url := fmt.Sprintf("http://127.0.0.1:%d/kick", n.apiPort)
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url,
		bytes.NewBuffer(usernamesByte))
	if err != nil {
		logrus.Errorf("Hysteria2 KickUsers NewRequest err: %v", err)
		return errors.New(constant.SysError)
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := http.DefaultClient.Do(req)
	defer func() {
		if resp != nil {
			resp.Body.Close()
		}
	}()
	if err != nil || resp.StatusCode != http.StatusOK {
		logrus.Errorf("Hysteria2 KickUsers err: %v", err)
		return errors.New("http connection error")
	}
	return nil
}

// OnlineUsers 在线用户
func (n *Hysteria2Api) OnlineUsers() (map[string]int64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	url := fmt.Sprintf("http://127.0.0.1:%d/online", n.apiPort)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		logrus.Errorf("Hysteria2 OnlineUsers NewRequest err: %v", err)
		return nil, errors.New(constant.SysError)
	}
	resp, err := http.DefaultClient.Do(req)
	defer func() {
		if resp != nil {
			_ = resp.Body.Close()
		}
	}()
	if err != nil || resp.StatusCode != http.StatusOK {
		logrus.Errorf("Hysteria2 OnlineUsers err: %v", err)
		return nil, errors.New("http connection error")
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		logrus.Errorf("Hysteria2 io read err: %v", err)
		return nil, errors.New("http connection error")
	}
	var onlineUsers map[string]int64
	if err = json.Unmarshal(body, &onlineUsers); err != nil {
		logrus.Errorf("Hysteria2 OnlineUsers Unmarshal err: %v", err)
		return nil, errors.New(constant.SysError)
	}
	return onlineUsers, nil
}

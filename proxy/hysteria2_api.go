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
	"os/exec"
	"time"
)

type Hysteria2Api struct {
	apiPort int64
	cmd     *exec.Cmd
}

func NewHysteria2Api(apiPort int64) *Hysteria2Api {
	return &Hysteria2Api{
		apiPort: apiPort,
		cmd:     &cmdHysteria2,
	}
}

func (h *Hysteria2Api) IsRunning() bool {
	return h.cmd != nil && h.cmd.Process != nil
}

// ListUsers 每个用户的流量信息
func (h *Hysteria2Api) ListUsers(clear bool) (map[string]bo.Hysteria2UserTraffic, error) {
	var users map[string]bo.Hysteria2UserTraffic
	if !h.IsRunning() {
		return users, nil
	}
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	url := fmt.Sprintf("http://127.0.0.1:%d/traffic", h.apiPort)
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
	if err = json.Unmarshal(body, &users); err != nil {
		logrus.Errorf("Hysteria2 ListUsers Unmarshal err: %v", err)
		return nil, errors.New(constant.SysError)
	}
	return users, nil
}

// KickUsers 踢下线
func (h *Hysteria2Api) KickUsers(keys []string) error {
	if !h.IsRunning() {
		return nil
	}
	usernamesByte, err := json.Marshal(keys)
	if err != nil {
		logrus.Errorf("Hysteria2 KickUsers Marshal err: %v", err)
		return errors.New(constant.SysError)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	url := fmt.Sprintf("http://127.0.0.1:%d/kick", h.apiPort)
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
func (h *Hysteria2Api) OnlineUsers() (map[string]int64, error) {
	var onlineUsers map[string]int64
	if !h.IsRunning() {
		return onlineUsers, nil
	}
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	url := fmt.Sprintf("http://127.0.0.1:%d/online", h.apiPort)
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
	if err = json.Unmarshal(body, &onlineUsers); err != nil {
		logrus.Errorf("Hysteria2 OnlineUsers Unmarshal err: %v", err)
		return nil, errors.New(constant.SysError)
	}
	return onlineUsers, nil
}

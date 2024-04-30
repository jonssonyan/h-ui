package proxy

import (
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

func apiClient() *http.Client {
	return &http.Client{
		Timeout: 5 * time.Second,
	}
}

func (n *Hysteria2Api) ListUsers(clear bool) (map[string]bo.Hysteria2UserTraffic, error) {
	client := apiClient()
	url := fmt.Sprintf("http://127.0.0.1:%d/traffic", n.apiPort)
	if clear {
		url = fmt.Sprintf("%s?clear=1", url)
	}
	resp, err := client.Get(url)
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

package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"h-ui/dao"
	"h-ui/model/constant"
	"h-ui/util"
	"net/http"
	"strconv"
)

type WebServer struct {
	server  http.Server
	ctx     context.Context
	cancel  context.CancelFunc
	newPort int
}

var webServer *WebServer

func NewServer(r *gin.Engine) (*WebServer, error) {
	if webServer != nil {
		return webServer, nil
	}
	config, err := dao.GetConfig("key = ?", constant.HUIWebPort)
	if err != nil {
		logrus.Errorf(err.Error())
		return nil, errors.New(err.Error())
	}

	newPort, err := strconv.Atoi(*config.Value)
	if err != nil {
		logrus.Errorf("conv newPort err: %v", err)
		return nil, errors.New(fmt.Sprintf("conv newPort err: %v", err))
	}
	if !util.IsPortAvailable(uint(newPort), "tcp") {
		logrus.Errorf("port is not available port: %d", newPort)
		return nil, errors.New(fmt.Sprintf("port is not available port: %d", newPort))
	}

	ctx, cancel := context.WithCancel(context.Background())

	webServer = &WebServer{
		server: http.Server{
			Addr:    fmt.Sprintf(":%d", newPort),
			Handler: r,
		},
		ctx:     ctx,
		cancel:  cancel,
		newPort: newPort,
	}
	return webServer, nil
}

func (w *WebServer) StartServer() error {
	crtPath, keyPath, err := w.getTLSConfigPaths()
	if err != nil {
		return err
	}
	if crtPath != "" && keyPath != "" {
		return w.server.ListenAndServeTLS(crtPath, keyPath)
	}
	return w.server.ListenAndServe()
}

func (w *WebServer) StopServer() error {
	if err := StopHysteria2(); err != nil {
		logrus.Errorf(err.Error())
		return err
	}

	if err := w.server.Shutdown(w.ctx); err != nil {
		logrus.Errorf("failed to shutdown server: %v", err)
		return err
	}
	w.cancel()
	webServer = nil
	return nil
}

func (w *WebServer) getTLSConfigPaths() (string, string, error) {
	crtPathConfig, err := dao.GetConfig("key = ?", constant.HUICrtPath)
	if err != nil {
		logrus.Errorf(err.Error())
		return "", "", errors.New(err.Error())
	}
	keyPathConfig, err := dao.GetConfig("key = ?", constant.HUIKeyPath)
	if err != nil {
		logrus.Errorf(err.Error())
		return "", "", errors.New(err.Error())
	}

	crtPath := ""
	if crtPathConfig.Value != nil && *crtPathConfig.Value != "" && util.Exists(*crtPathConfig.Value) {
		crtPath = *crtPathConfig.Value
	}

	keyPath := ""
	if keyPathConfig.Value != nil && *keyPathConfig.Value != "" && util.Exists(*keyPathConfig.Value) {
		keyPath = *keyPathConfig.Value
	}

	return crtPath, keyPath, nil
}

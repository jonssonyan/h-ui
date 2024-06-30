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
	server    http.Server
	ctx       context.Context
	cancel    context.CancelFunc
	httpPort  int
	HttpsPort int
}

var webServer *WebServer

func NewServer() (*WebServer, error) {
	if webServer != nil {
		return webServer, nil
	}
	httpPort, err := getPort(constant.HUIWebPort)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithCancel(context.Background())

	webServer = &WebServer{
		server: http.Server{
			Addr: fmt.Sprintf(":%d", httpPort),
		},
		ctx:      ctx,
		cancel:   cancel,
		httpPort: httpPort,
	}
	return webServer, nil
}

func getPort(configKey string) (int, error) {
	config, err := dao.GetConfig("key = ?", configKey)
	if err != nil {
		return 0, errors.New(err.Error())
	}

	port, err := strconv.Atoi(*config.Value)
	if err != nil {
		logrus.Errorf("conv port err: %v", err)
		return 0, errors.New(fmt.Sprintf("port: %s is invalid", *config.Value))
	}

	if !util.IsPortAvailable(uint(port), "tcp") {
		logrus.Errorf("port is not available port: %d", port)
		return 0, errors.New(fmt.Sprintf("port is not available port: %d", port))
	}

	return port, nil
}

func (w *WebServer) StartServer(r *gin.Engine) error {
	w.server.Handler = r
	w.HttpsPort = 0

	httpsPort, err := getPort(constant.HUIWebHttpsPort)
	if err != nil {
		return err
	}

	crtPath, keyPath, err := w.getTLSConfigPaths()
	if err != nil {
		return err
	}

	if httpsPort > 0 && crtPath != "" && keyPath != "" {
		w.HttpsPort = httpsPort
		go func() {
			w.server.Addr = fmt.Sprintf(":%d", w.HttpsPort)
			if err := w.server.ListenAndServeTLS(crtPath, keyPath); err != nil {
				logrus.Errorf("failed to start ListenAndServeTLS: %v", err)
			}
		}()
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

func (w *WebServer) IsHttps() bool {
	return w.HttpsPort > 0
}

package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"h-ui/util"
	"net/http"
)

type WebServer struct {
	ctx    context.Context
	cancel context.CancelFunc
	server *http.Server
}

var webServer *WebServer

func NewServer() (*WebServer, error) {
	if webServer != nil {
		return webServer, nil
	}

	ctx, cancel := context.WithCancel(context.Background())
	webServer = &WebServer{
		ctx:    ctx,
		cancel: cancel,
	}
	return webServer, nil
}

func (w *WebServer) StartServer(r *gin.Engine) error {
	port, crtPath, keyPath, err := w.getPortAndCert()
	if err != nil {
		return err
	}

	w.server = &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: r,
	}

	if crtPath != "" && keyPath != "" {
		return w.server.ListenAndServeTLS(crtPath, keyPath)
	}
	return w.server.ListenAndServe()
}

func (w *WebServer) StopServer() error {
	if err := StopHysteria2(); err != nil {
		return err
	}

	defer w.cancel()

	if w.server != nil {
		if err := w.server.Shutdown(w.ctx); err != nil {
			logrus.Errorf("failed to shutdown server: %v", err)
			return err
		}
	}

	webServer = nil
	return nil
}

func (w *WebServer) getPortAndCert() (int64, string, string, error) {
	port, crtPath, keyPath, err := GetPortAndCert()
	if err != nil {
		return 0, "", "", err
	}

	if !util.IsPortAvailable(uint(port), "tcp") {
		errMsg := fmt.Sprintf("port is not available: %d", port)
		logrus.Error(errMsg)
		return 0, "", "", errors.New(errMsg)
	}

	if crtPath != "" && !util.Exists(crtPath) {
		errMsg := fmt.Sprintf("crt path: %s does not exist", crtPath)
		logrus.Error(errMsg)
		return 0, "", "", errors.New(errMsg)
	}

	if keyPath != "" && !util.Exists(keyPath) {
		errMsg := fmt.Sprintf("key path: %s does not exist", keyPath)
		logrus.Error(errMsg)
		return 0, "", "", errors.New(errMsg)
	}

	return port, crtPath, keyPath, nil
}

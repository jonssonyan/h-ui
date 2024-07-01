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

	if err := w.server.Shutdown(w.ctx); err != nil {
		logrus.Errorf("failed to shutdown server: %v", err)
		return err
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
		logrus.Errorf("port is not available port: %d", port)
		return 0, "", "", errors.New(fmt.Sprintf("port is not available port: %d", port))
	}

	if crtPath != "" {
		if !util.Exists(crtPath) {
			logrus.Errorf("crt path: %s is not exist", crtPath)
			return 0, "", "", errors.New(fmt.Sprintf("crt path: %s is not exist", crtPath))
		}
	}

	if keyPath != "" {
		if !util.Exists(keyPath) {
			logrus.Errorf("key path: %s is not exist", keyPath)
			return 0, "", "", errors.New(fmt.Sprintf("key path: %s is not exist", keyPath))
		}
	}

	return port, crtPath, keyPath, nil
}

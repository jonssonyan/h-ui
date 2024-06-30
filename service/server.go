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

	httpServer  http.Server
	httpsServer http.Server
	httpPort    int64
	httpsPort   int64
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
	httpPort, httpsPort, crtPath, keyPath, err := w.getPortAndCert()
	if err != nil {
		return err
	}

	if httpsPort > 0 && crtPath != "" && keyPath != "" {
		go func() {
			w.httpsServer.Handler = r
			w.httpsServer.Addr = fmt.Sprintf(":%d", httpsPort)
			w.httpsPort = httpsPort
			if err := w.httpsServer.ListenAndServeTLS(crtPath, keyPath); err != nil {
				logrus.Errorf("failed to start ListenAndServeTLS: %v", err)
			}
		}()
	}
	w.httpServer.Handler = r
	w.httpServer.Addr = fmt.Sprintf(":%d", httpPort)
	w.httpPort = httpPort
	return w.httpServer.ListenAndServe()
}

func (w *WebServer) StopServer() error {
	if err := StopHysteria2(); err != nil {
		logrus.Errorf(err.Error())
		return err
	}

	if err := w.httpServer.Shutdown(w.ctx); err != nil {
		logrus.Errorf("failed to shutdown server: %v", err)
		return err
	}

	if w.httpsPort > 0 {
		if err := w.httpsServer.Shutdown(w.ctx); err != nil {
			logrus.Errorf("failed to shutdown https server: %v", err)
			return err
		}
	}

	w.cancel()
	webServer = nil
	return nil
}

func (w *WebServer) getPortAndCert() (int64, int64, string, string, error) {
	httpPort, httpsPort, crtPath, keyPath, err := GetPortAndCert()
	if err != nil {
		return 0, 0, "", "", err
	}

	if !util.IsPortAvailable(uint(httpPort), "tcp") {
		logrus.Errorf("port is not available port: %d", httpPort)
		return 0, 0, "", "", errors.New(fmt.Sprintf("port is not available port: %d", httpPort))
	}

	if !util.IsPortAvailable(uint(httpsPort), "tcp") {
		logrus.Errorf("https port is not available port: %d", httpsPort)
		return 0, 0, "", "", errors.New(fmt.Sprintf("https port is not available port: %d", httpsPort))
	}

	if httpsPort > 0 {
		if !util.Exists(crtPath) {
			return 0, 0, "", "", errors.New(fmt.Sprintf("crt path: %s is not exist", crtPath))
		}

		if !util.Exists(keyPath) {
			return 0, 0, "", "", errors.New(fmt.Sprintf("key path: %s is not exist", keyPath))
		}
	}

	return httpPort, httpsPort, crtPath, keyPath, nil
}

func (w *WebServer) GetHttps() int64 {
	return w.httpsPort
}

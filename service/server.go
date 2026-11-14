package service

import (
	"context"
	"crypto/tls"
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	"h-ui/util"
	"net/http"
	"time"
)

var server *http.Server

func InitServer(addr string, handler http.Handler, certPath, keyPath string) {
	server = &http.Server{
		Addr:    addr,
		Handler: handler,
	}

	if certPath != "" && keyPath != "" {
		server.TLSConfig = &tls.Config{
			// GetCertificate will be called in every TLS handshake, So the
			// server will always get the latest certificate.
			GetCertificate: func(*tls.ClientHelloInfo) (*tls.Certificate, error) {
				cert, err := tls.LoadX509KeyPair(certPath, keyPath)

				if err != nil {
					return nil, err
				}

				return &cert, nil
			},
		}

	}
}

func StartServer() error {
	if server.TLSConfig != nil {
		return server.ListenAndServeTLS("", "")
	}
	return server.ListenAndServe()
}

func StopServer() error {
	if err := StopHysteria2(); err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		logrus.Errorf("failed to shutdown server: %v", err)
		return errors.New("failed to shutdown server")
	}

	return nil
}

func GetServerPortAndCert() (int64, string, string, error) {
	port, crtPath, keyPath, err := GetPortAndCert()
	if err != nil {
		return 0, "", "", err
	}

	if !util.IsPortAvailable(uint(port), "tcp") {
		errMsg := fmt.Sprintf("port %d is taken", port)
		logrus.Errorf(errMsg)
		return 0, "", "", errors.New(errMsg)
	}

	if crtPath != "" && !util.Exists(crtPath) {
		errMsg := fmt.Sprintf("crt path: %s does not exist", crtPath)
		logrus.Errorf(errMsg)
		return 0, "", "", errors.New(errMsg)
	}

	if keyPath != "" && !util.Exists(keyPath) {
		errMsg := fmt.Sprintf("key path: %s does not exist", keyPath)
		logrus.Errorf(errMsg)
		return 0, "", "", errors.New(errMsg)
	}

	return port, crtPath, keyPath, nil
}

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

var (
	server *http.Server
	ctx    context.Context
	cancel context.CancelFunc
)

func StartServer(r *gin.Engine) error {
	config, err := dao.GetConfig("key = ?", constant.HUIWebPort)
	if err != nil {
		logrus.Errorf(err.Error())
		return errors.New(err.Error())
	}

	newPortStr := *config.Value
	newPort, err := strconv.Atoi(newPortStr)
	if err != nil {
		logrus.Errorf("conv newPort err: %v", err)
		return errors.New(fmt.Sprintf("conv newPort err: %v", err))
	}
	available := util.IsPortAvailable(uint(newPort), "tcp")
	if !available {
		logrus.Errorf("port is not available port: %d", newPort)
		return errors.New(fmt.Sprintf("port is not available port: %d", newPort))
	}

	server = &http.Server{
		Addr:    fmt.Sprintf(":%d", newPort),
		Handler: r,
	}
	ctx, cancel = context.WithCancel(context.Background())
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		logrus.Errorf("could not listen on port %d: %v", newPort, err)
		return err
	}

	logrus.Infof("server is running on port %d", newPort)
	return nil
}

func StopServer() error {
	if server != nil {
		cancel()
		if err := server.Shutdown(ctx); err != nil {
			logrus.Errorf("failed to shutdown server: %v", err)
			return err
		}
	}
	return nil
}

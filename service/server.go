package service

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"h-ui/dao"
	"h-ui/model/constant"
	"h-ui/util"
	"net/http"
	"strconv"
	"sync"
)

var (
	server *http.Server
	mutex  sync.Mutex
)

func StartServer(r *gin.Engine) error {
	if !mutex.TryLock() {
		return errors.New("start server err")
	}
	defer mutex.Unlock()

	config, err := dao.GetConfig("key = ?", constant.HUIWebPort)
	if err != nil {
		return errors.New(err.Error())
	}

	newPortStr := *config.Value
	newPort, err := strconv.Atoi(newPortStr)
	if err != nil {
		return errors.New(fmt.Sprintf("conv newPort err: %v", err))
	}
	available := util.IsPortAvailable(uint(newPort), "tcp")
	if !available {
		return errors.New(fmt.Sprintf("port is not available port: %d", newPort))
	}

	if server != nil {
		if err := server.Shutdown(nil); err != nil {
			logrus.Errorf("failed to shutdown server: %v", err)
		}
	}

	server = &http.Server{
		Addr:    fmt.Sprintf(":%d", newPort),
		Handler: r,
	}

	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		logrus.Errorf("could not listen on port %d: %v", newPort, err)
	}

	logrus.Infof("server is running on port %d", newPort)
	return nil
}

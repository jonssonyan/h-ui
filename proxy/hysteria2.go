package proxy

import (
	"context"
	"errors"
	"github.com/sirupsen/logrus"
	"h-ui/dao"
	"h-ui/model/constant"
	"h-ui/util"
	"os/exec"
	"sync"
	"time"
)

var mutexHysteria2 sync.Mutex
var cmdMapHysteria2 sync.Map

type Hysteria2Process struct {
	process
}

func NewHysteria2Instance() *Hysteria2Process {
	return &Hysteria2Process{process{mutex: &mutexHysteria2, cmdMap: &cmdMapHysteria2}}
}

func (h *Hysteria2Process) StopHysteria2Instance() error {
	config, err := dao.GetConfig(constant.Hysteria2Config)
	if err != nil {
		return err
	}
	if err = h.Stop(*config.Value, true); err != nil {
		return err
	}
	return nil
}

func (h *Hysteria2Process) StartHysteria2(port string) error {
	defer h.mutex.Unlock()
	if h.mutex.TryLock() {
		if h.IsRunning(port) {
			return nil
		}
		binaryFilePath := util.GetHysteria2BinPath()
		configFilePath := constant.Hysteria2ConfigPath
		cmd := exec.Command(binaryFilePath, "-c", configFilePath, "server")
		if cmd.Err != nil {
			if err := util.RemoveFile(configFilePath); err != nil {
				return err
			}
			logrus.Errorf("hysteria2 command error err: %v", cmd.Err)
			return errors.New(constant.Hysteria2StartError)
		}
		if err := cmd.Start(); err != nil {
			if err = util.RemoveFile(configFilePath); err != nil {
				return err
			}
			logrus.Errorf("start hysteria2 error err: %v", err)
			return errors.New(constant.Hysteria2StartError)
		}
		h.cmdMap.Store(port, cmd)

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		done := make(chan error)
		go func() {
			done <- cmd.Wait()
			select {
			case err := <-done:
				if err != nil {
					logrus.Errorf("hysteria2 process wait error err: %v", err)
					h.releaseProcess(port)
				}
			case <-ctx.Done():
				logrus.Errorf("hysteria2 process wait timeout")
				h.releaseProcess(port)
			}
		}()
		return nil
	}
	logrus.Errorf("start hysteria2 error err: lock not acquired")
	return errors.New(constant.Hysteria2StartError)
}

func (h *Hysteria2Process) releaseProcess(port string) {
	load, ok := NewHysteria2Instance().GetCmdMap().Load(port)
	if ok {
		cmd := load.(*exec.Cmd)
		if !cmd.ProcessState.Success() {
			h.cmdMap.Delete(port)
			if err := cmd.Process.Release(); err != nil {
				logrus.Errorf("hysteria2 process release error err: %v", err)
			}
			if err := util.RemoveFile(constant.Hysteria2ConfigPath); err != nil {
				logrus.Errorf("hysteria2 process remove file error err: %v", err)
			}
		}
	}
}

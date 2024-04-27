package proxy

import (
	"errors"
	"github.com/sirupsen/logrus"
	"h-ui/model/constant"
	"h-ui/util"
	"os/exec"
	"sync"
)

type process struct {
	mutex  *sync.Mutex
	cmdMap *sync.Map
}

func (p *process) GetCmdMap() *sync.Map {
	return p.cmdMap
}

func (p *process) IsRunning(port string) bool {
	cmd, ok := p.cmdMap.Load(port)
	if ok {
		if cmd == nil || cmd.(*exec.Cmd).Process == nil {
			return false
		}
		if cmd.(*exec.Cmd).ProcessState == nil {
			return true
		}
	}
	return false
}

func (p *process) Stop(port string, removeFile bool) error {
	defer p.mutex.Unlock()
	if p.mutex.TryLock() {
		if !p.IsRunning(port) {
			logrus.Errorf("process has been stoped. port: %s", port)
			if removeFile {
				if err := util.RemoveFile(constant.Hysteria2ConfigPath); err != nil {
					return err
				}
			}
			return nil
		}
		cmd, ok := p.cmdMap.Load(port)
		if ok {
			if err := cmd.(*exec.Cmd).Process.Kill(); err != nil {
				logrus.Errorf("stop process error. port: %s err: %v", port, err)
				return errors.New(constant.ProcessStopError)
			}
			p.cmdMap.Delete(port)
			if removeFile {
				if err := util.RemoveFile(constant.Hysteria2ConfigPath); err != nil {
					return err
				}
			}
			return nil
		}
		logrus.Errorf("stop process error port: %s err: process not found", port)
		return errors.New(constant.ProcessStopError)
	}
	logrus.Errorf("stop process error err: lock not acquired")
	return errors.New(constant.ProcessStopError)
}

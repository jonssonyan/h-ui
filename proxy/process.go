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

func (p *process) IsRunning(apiPort string) bool {
	cmd, ok := p.cmdMap.Load(apiPort)
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

func (p *process) Stop(apiPort string, removeFile bool) error {
	defer p.mutex.Unlock()
	if p.mutex.TryLock() {
		if !p.IsRunning(apiPort) {
			logrus.Errorf("process has been stoped. apiPort: %s", apiPort)
			if removeFile {
				if err := util.RemoveFile(constant.Hysteria2ConfigPath); err != nil {
					return err
				}
			}
			return nil
		}
		cmd, ok := p.cmdMap.Load(apiPort)
		if ok {
			if err := cmd.(*exec.Cmd).Process.Kill(); err != nil {
				logrus.Errorf("stop process error. apiPort: %s err: %v", apiPort, err)
				return errors.New(constant.ProcessStopError)
			}
			p.cmdMap.Delete(apiPort)
			if removeFile {
				if err := util.RemoveFile(constant.Hysteria2ConfigPath); err != nil {
					return err
				}
			}
			return nil
		}
		logrus.Errorf("stop process error apiPort: %s err: process not found", apiPort)
		return errors.New(constant.ProcessStopError)
	}
	logrus.Errorf("stop process error err: lock not acquired")
	return errors.New(constant.ProcessStopError)
}

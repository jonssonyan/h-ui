package proxy

import (
	"context"
	"errors"
	"github.com/sirupsen/logrus"
	"os/exec"
	"sync"
	"time"
)

type process struct {
	mutex *sync.Mutex
	cmd   *exec.Cmd
}

func (p *process) IsRunning() bool {
	return p.cmd != nil && p.cmd.Process != nil
}

func (p *process) Start(name string, arg ...string) error {
	defer p.mutex.Unlock()
	if p.mutex.TryLock() {
		if p.IsRunning() {
			return nil
		}

		cmd := exec.Command(name, arg...)
		if cmd.Err != nil {
			logrus.Errorf("cmd err: %v", cmd.Err)
			return errors.New("cmd err")
		}
		if err := cmd.Start(); err != nil {
			logrus.Errorf("cmd start err: %v", err)
			return errors.New("cmd start err")
		}

		p.cmd = cmd

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		done := make(chan error)
		go func() {
			done <- cmd.Wait()
			select {
			case err := <-done:
				if err != nil {
					logrus.Errorf("cmd wait err: %v", err)
					_ = p.releaseProcess()
				}
			case <-ctx.Done():
				logrus.Errorf("cmd wait timeout")
				_ = p.releaseProcess()
			}
		}()
		return nil
	}
	logrus.Errorf("start cmd err: lock not acquired")
	return errors.New("start cmd err")
}

func (p *process) Stop() error {
	defer p.mutex.Unlock()
	if p.mutex.TryLock() {
		if !p.IsRunning() {
			return nil
		}
		if err := p.cmd.Process.Kill(); err != nil {
			logrus.Errorf("cmd stop err: %v", err)
			return errors.New("cmd stop err")
		}
		p.cmd = nil
		return nil
	}
	logrus.Errorf("cmd stop err: lock not acquired")
	return errors.New("cmd stop err")
}

func (p *process) releaseProcess() error {
	if p.cmd != nil && p.cmd.Process != nil {
		if err := p.cmd.Process.Release(); err != nil {
			logrus.Errorf("cmd release err: %v", err)
			return errors.New("cmd release err")
		}
		p.cmd = nil
	}
	return nil
}

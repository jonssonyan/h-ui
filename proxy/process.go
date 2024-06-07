package proxy

import (
	"bufio"
	"context"
	"errors"
	"github.com/sirupsen/logrus"
	"h-ui/model/constant"
	"io"
	"os/exec"
	"sync"
	"time"
)

var loggerPlus *LoggerPlus

func init() {
	loggerPlus = NewLoggerPlus(constant.Hysteria2LogPath, 1, 2, 30, true)
}

type process struct {
	mutex *sync.Mutex
	cmd   *exec.Cmd
}

func (p *process) IsRunning() bool {
	return p.cmd != nil && p.cmd.Process != nil
}

func (p *process) Start(name string, arg ...string) error {
	if !p.mutex.TryLock() {
		logrus.Errorf("start cmd err: lock not acquired")
		return errors.New("start cmd err")
	}
	defer p.mutex.Unlock()

	if p.IsRunning() {
		return nil
	}

	cmd := exec.Command(name, arg...)
	if cmd.Err != nil {
		logrus.Errorf("cmd err: %v", cmd.Err)
		return errors.New("cmd err")
	}

	// 获取命令的 stdout 和 stderr
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		logrus.Errorf("Error obtaining stdout: %v", err)
		return err
	}
	stderr, err := cmd.StderrPipe()
	if err != nil {
		logrus.Errorf("Error obtaining stderr: %v", err)
		return err
	}

	if err := cmd.Start(); err != nil {
		logrus.Errorf("cmd start err: %v", err)
		return errors.New("cmd start err")
	}

	p.cmd = cmd

	go p.handleCmdExecution(cmd)

	go p.handleLogs(stdout, stderr)

	return nil
}

func (p *process) Stop() error {
	if !p.mutex.TryLock() {
		logrus.Errorf("cmd stop err: lock not acquired")
		return errors.New("cmd stop err")
	}
	defer p.mutex.Unlock()

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

func (p *process) release() error {
	if !p.mutex.TryLock() {
		logrus.Errorf("cmd release err: lock not acquired")
		return errors.New("cmd release err")
	}
	defer p.mutex.Unlock()

	if !p.IsRunning() {
		return nil
	}

	if err := p.cmd.Process.Release(); err != nil {
		logrus.Errorf("cmd release err: %v", err)
		return errors.New("cmd release err")
	}
	p.cmd = nil
	return nil
}

func (p *process) handleCmdExecution(cmd *exec.Cmd) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	done := make(chan error, 1)
	go func() {
		done <- cmd.Wait()
	}()

	select {
	case err := <-done:
		if err != nil {
			logrus.Errorf("cmd wait err: %v", err)
			_ = p.release()
		}
	case <-ctx.Done():
		logrus.Errorf("cmd wait timeout")
		_ = p.release()
	}
}

func (p *process) handleLogs(stdout, stderr io.ReadCloser) {
	// 日志
	stdoutChan := make(chan string)
	stderrChan := make(chan string)

	go func() {
		scanner := bufio.NewScanner(stdout)
		for scanner.Scan() {
			stdoutChan <- scanner.Text()
		}
		if err := scanner.Err(); err != nil {
			logrus.Errorf("Error reading stdout: %v", err)
		}
		close(stdoutChan)
	}()

	go func() {
		scanner := bufio.NewScanner(stderr)
		for scanner.Scan() {
			stderrChan <- scanner.Text()
		}
		if err := scanner.Err(); err != nil {
			logrus.Errorf("Error reading stderr: %v", err)
		}
		close(stderrChan)
	}()

	for {
		select {
		case line, ok := <-stdoutChan:
			if !ok {
				stdoutChan = nil
			} else {
				loggerPlus.Infof(line)
			}
		case line, ok := <-stderrChan:
			if !ok {
				stderrChan = nil
			} else {
				loggerPlus.Errorf(line)
			}
		}

		// 当两个 channel 都关闭时，退出循环
		if stdoutChan == nil && stderrChan == nil {
			break
		}
	}
}

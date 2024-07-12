package proxy

import (
	"bufio"
	"errors"
	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
	"h-ui/model/constant"
	"io"
	"os/exec"
	"sync"
)

var logger logrus.Logger

func initLogger() {
	logger.SetOutput(&lumberjack.Logger{
		Filename:   constant.Hysteria2LogPath,
		MaxSize:    1,
		MaxBackups: 2,
		MaxAge:     30,
		Compress:   true,
		LocalTime:  true,
	})
	logger.SetFormatter(&logrus.JSONFormatter{TimestampFormat: "2006-01-02 15:04:05"})
	logger.SetLevel(logrus.WarnLevel)
}

func init() {
	initLogger()
}

type process struct {
	mutex *sync.Mutex
	cmd   *exec.Cmd
}

func (p *process) isRunning() bool {
	return p.cmd != nil && p.cmd.Process != nil && p.cmd.ProcessState == nil
}

func (p *process) start(name string, arg ...string) error {
	if !p.mutex.TryLock() {
		logrus.Errorf("start cmd err: lock not acquired")
		return errors.New("start cmd err")
	}
	defer p.mutex.Unlock()

	if p.isRunning() {
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

	go p.handleLogs(stdout, stderr)

	return nil
}

func (p *process) stop() error {
	if !p.mutex.TryLock() {
		logrus.Errorf("cmd stop err: lock not acquired")
		return errors.New("cmd stop err")
	}
	defer p.mutex.Unlock()

	if !p.isRunning() {
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

	if !p.isRunning() {
		return nil
	}

	if err := p.cmd.Process.Release(); err != nil {
		logrus.Errorf("cmd release err: %v", err)
		return errors.New("cmd release err")
	}
	p.cmd = nil
	return nil
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
				logger.Infof(line)
			}
		case line, ok := <-stderrChan:
			if !ok {
				stderrChan = nil
			} else {
				logger.Errorf(line)
			}
		}

		// 当两个 channel 都关闭时，退出循环
		if stdoutChan == nil && stderrChan == nil {
			break
		}
	}
}

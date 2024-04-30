package proxy

import (
	"h-ui/util"
	"os/exec"
	"sync"
)

var mutexHysteria2 sync.Mutex
var cmdHysteria2 exec.Cmd

type Hysteria2Process struct {
	process
	port       string
	binPath    string
	configPath string
}

func NewHysteria2Instance(port string, binPath string, configPath string) *Hysteria2Process {
	return &Hysteria2Process{process{mutex: &mutexHysteria2, cmd: &cmdHysteria2}, port, binPath, configPath}
}

func (h *Hysteria2Process) StartHysteria2() error {
	if err := h.Start(h.binPath, "-c", h.configPath, "server"); err != nil {
		_ = util.RemoveFile(h.configPath)
		return err
	}
	return nil
}

func (h *Hysteria2Process) StopHysteria2() error {
	if err := h.Stop(); err != nil {
		return err
	}
	_ = util.RemoveFile(h.configPath)
	return nil
}

func (h *Hysteria2Process) RestartHysteria2() error {
	if err := h.Stop(); err != nil {
		return err
	}
	if err := h.StartHysteria2(); err != nil {
		return err
	}
	return nil
}

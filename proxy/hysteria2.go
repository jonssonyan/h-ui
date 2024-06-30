package proxy

import (
	"h-ui/model/constant"
	"h-ui/util"
	"os/exec"
	"sync"
)

type Hysteria2Process struct {
	process
	binPath    string
	configPath string
}

var mutexHysteria2 sync.Mutex
var cmdHysteria2 exec.Cmd
var hysteria2Instance *Hysteria2Process

func init() {
	hysteria2Instance = &Hysteria2Process{process{mutex: &mutexHysteria2, cmd: &cmdHysteria2}, util.GetHysteria2BinPath(), constant.Hysteria2ConfigPath}
}

func NewHysteria2Instance() *Hysteria2Process {
	return hysteria2Instance
}

func (h *Hysteria2Process) IsRunning() bool {
	return h.isRunning()
}

func (h *Hysteria2Process) StartHysteria2() error {
	if err := h.start(h.binPath, "-c", h.configPath, "server"); err != nil {
		_ = util.RemoveFile(h.configPath)
		return err
	}
	return nil
}

func (h *Hysteria2Process) StopHysteria2() error {
	if err := h.stop(); err != nil {
		return err
	}
	_ = util.RemoveFile(h.configPath)
	return nil
}

func (h *Hysteria2Process) Release() error {
	return h.release()
}

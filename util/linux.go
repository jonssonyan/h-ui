package util

import (
	"fmt"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/mem"
	"github.com/sirupsen/logrus"
	"net"
	"os"
	"os/exec"
	"strconv"
	"time"
)

func Exec(cmd string) (string, error) {
	command := exec.Command("bash", "-c", cmd)
	command.Env = os.Environ()
	output, err := command.CombinedOutput()
	if err != nil {
		logrus.Errorf("execute command failed cmd: %s err: %v", cmd, err)
		return "", fmt.Errorf("execute command failed cmd: %s", cmd)
	}
	return string(output), nil
}

func IsPortAvailable(port uint, network string) bool {
	if network == "tcp" {
		listener, err := net.ListenTCP(network, &net.TCPAddr{
			IP:   net.IPv4(0, 0, 0, 0),
			Port: int(port),
		})
		defer func() {
			if listener != nil {
				listener.Close()
			}
		}()
		if err != nil {
			logrus.Errorf("port %d is taken err: %s", port, err)
			return false
		}
	}
	if network == "udp" {
		listener, err := net.ListenUDP("udp", &net.UDPAddr{
			IP:   net.IPv4(0, 0, 0, 0),
			Port: int(port),
		})
		defer func() {
			if listener != nil {
				listener.Close()
			}
		}()
		if err != nil {
			logrus.Errorf("port %d is taken err: %s", port, err)
			return false
		}
	}
	return true
}

func GetCpuPercent() (float64, error) {
	var err error
	percent, err := cpu.Percent(time.Second, false)
	value, err := strconv.ParseFloat(fmt.Sprintf("%.1f", percent[0]), 64)
	return value, err
}

func GetMemPercent() (float64, error) {
	var err error
	memInfo, err := mem.VirtualMemory()
	value, err := strconv.ParseFloat(fmt.Sprintf("%.1f", memInfo.UsedPercent), 64)
	return value, err
}

func GetDiskPercent() (float64, error) {
	var err error
	parts, err := disk.Partitions(true)
	diskInfo, err := disk.Usage(parts[0].Mountpoint)
	value, err := strconv.ParseFloat(fmt.Sprintf("%.1f", diskInfo.UsedPercent), 64)
	return value, err
}

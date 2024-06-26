package util

import (
	"fmt"
	"h-ui/model/constant"
	"io"
	"net/http"
	"os"
	"runtime"
)

func GetHysteria2BinPath() string {
	return constant.BinDir + GetHysteria2BinName()
}

func GetHysteria2BinName() string {
	hysteria2FileName := fmt.Sprintf("hysteria-%s-%s", runtime.GOOS, runtime.GOARCH)
	if runtime.GOOS == "windows" {
		hysteria2FileName += ".exe"
	}
	return hysteria2FileName
}

func DownloadHysteria2(version string) error {
	hysteria2BinName := GetHysteria2BinName()
	hysteria2BinPath := GetHysteria2BinPath()

	// Download the latest version of Hysteria2
	url, err := GetReleaseAssetURL("apernet", "hysteria", version, hysteria2BinName)
	if err != nil {
		return err
	}

	resp, err := http.Get(url)
	defer resp.Body.Close()
	if err != nil {
		return fmt.Errorf("failed to download file: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to download file, status code: %d", resp.StatusCode)
	}

	if Exists(hysteria2BinPath) {
		if err = os.Remove(hysteria2BinPath); err != nil {
			return fmt.Errorf("failed to remove existing file: %v", err)
		}
	}

	file, err := os.Create(hysteria2BinPath)
	defer file.Close()
	if err != nil {
		return fmt.Errorf("failed to create file %s: %v", hysteria2BinPath, err)
	}

	_, err = io.Copy(file, resp.Body)
	if err != nil {
		return fmt.Errorf("failed to write to file: %v", err)
	}

	if err = os.Chmod(hysteria2BinPath, 0755); err != nil {
		return fmt.Errorf("failed to change file permissions: %v", err)
	}
	return nil
}

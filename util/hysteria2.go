package util

import (
	"fmt"
	"h-ui/model/constant"
	"io"
	"net/http"
	"os"
	"runtime"
)

func GetHysteria2FileName() string {
	hysteria2FileName := fmt.Sprintf("hysteria-%s-%s", runtime.GOOS, runtime.GOARCH)
	if runtime.GOOS == "windows" {
		hysteria2FileName += ".exe"
	}
	return hysteria2FileName
}

func DownloadHysteria2(version string) error {
	url, err := GetReleaseAssetURL("apernet", "hysteria", version, GetHysteria2FileName())
	if err != nil {
		return err
	}
	hysteria2Path := constant.ProxyDir + GetHysteria2FileName()

	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("failed to download file: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to download file, status code: %d", resp.StatusCode)
	}

	if !Exists(hysteria2Path) {
		if err = os.Remove(hysteria2Path); err != nil {
			return fmt.Errorf("failed to remove existing file: %v", err)
		}
	}

	file, err := os.Create(hysteria2Path)
	if err != nil {
		return fmt.Errorf("failed to create file %s: %v", hysteria2Path, err)
	}
	defer file.Close()

	_, err = io.Copy(file, resp.Body)
	if err != nil {
		return fmt.Errorf("failed to write to file: %v", err)
	}

	if err = os.Chmod(hysteria2Path, 0755); err != nil {
		return fmt.Errorf("failed to change file permissions: %v", err)
	}
	return nil
}

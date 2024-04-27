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
	hysteria2FileName := fmt.Sprintf("hysteria2-%s-%s", runtime.GOOS, runtime.GOARCH)
	if runtime.GOOS == "windows" {
		hysteria2FileName += ".exe"
	}
	return constant.BinDir + hysteria2FileName
}

func DownloadHysteria2(version string) error {
	hysteria2Path := GetHysteria2BinPath()
	// 判断文件是否以及存在

	url, err := GetReleaseAssetURL("apernet", "hysteria", version, hysteria2Path)
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

	if Exists(hysteria2Path) {
		if err = os.Remove(hysteria2Path); err != nil {
			return fmt.Errorf("failed to remove existing file: %v", err)
		}
	}

	file, err := os.Create(hysteria2Path)
	defer file.Close()
	if err != nil {
		return fmt.Errorf("failed to create file %s: %v", hysteria2Path, err)
	}

	_, err = io.Copy(file, resp.Body)
	if err != nil {
		return fmt.Errorf("failed to write to file: %v", err)
	}

	if err = os.Chmod(hysteria2Path, 0755); err != nil {
		return fmt.Errorf("failed to change file permissions: %v", err)
	}
	return nil
}

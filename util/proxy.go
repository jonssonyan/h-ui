package util

import (
	"fmt"
	"h-ui/model/constant"
)

func GetProxyBin() (string, error) {
	hysteria2Path := constant.ProxyDir + GetHysteria2FileName()
	if !Exists(hysteria2Path) {
		if err := DownloadHysteria2(""); err != nil {
			return "", err
		}
	}
	return hysteria2Path, nil
}

func GetProxyConfig() (string, error) {
	hysteria2Path := fmt.Sprintf("%s%s.json", constant.ProxyDir, GetHysteria2FileName())
	if !Exists(hysteria2Path) {
		return "", fmt.Errorf("%s not exist", hysteria2Path)
	}
	return hysteria2Path, nil
}

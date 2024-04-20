package service

import (
	"h-ui/dao"
)

func GetConfig(key string) (string, error) {
	config, err := dao.GetConfig("key = ?", key)
	if err != nil {
		return "", err
	}
	return *config.Value, nil
}

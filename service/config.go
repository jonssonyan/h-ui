package service

import (
	"h-ui/dao"
	"h-ui/model/entity"
)

func UpdateConfig(key string, value string) error {
	return dao.UpdateConfig([]string{key}, map[string]interface{}{"value": value})
}

func ListConfig(keys []string) ([]entity.Config, error) {
	return dao.ListConfig("key in ?", keys)
}

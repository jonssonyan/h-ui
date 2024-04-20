package dao

import (
	"errors"
	"gorm.io/gorm"
	"h-ui/model/constant"
	"h-ui/model/entity"
	"time"
)

func SaveConfig(config entity.Config) (int64, error) {
	if tx := sqliteDB.Save(&config); tx.Error != nil {
		return 0, tx.Error
	}
	return *config.Id, nil
}

func UpdateConfig(ids []int64, updates map[string]interface{}) error {
	if len(updates) > 0 {
		updates["update_time"] = time.Now()
		if tx := sqliteDB.Model(&entity.Config{}).
			Where("id in ?", ids).
			Updates(updates); tx.Error != nil {
			return tx.Error
		}
	}
	return nil
}

func GetConfig(query interface{}, args ...interface{}) (entity.Config, error) {
	var config entity.Config
	if tx := sqliteDB.Model(&entity.Config{}).
		Where(query, args...).First(&config); tx.Error != nil {
		if tx.Error == gorm.ErrRecordNotFound {
			return config, errors.New(constant.ConfigNotExist)
		}
		return config, tx.Error
	}
	return config, nil
}

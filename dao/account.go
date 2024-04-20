package dao

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"h-ui/model/constant"
	"h-ui/model/dto"
	"h-ui/model/entity"
	"time"
)

func SaveAccount(account entity.Account) (int64, error) {
	if tx := sqliteDB.Save(&account); tx.Error != nil {
		return 0, tx.Error
	}
	return *account.Id, nil
}

func DeleteAccount(ids []int64) error {
	if tx := sqliteDB.Where("id in ?", ids).Delete(&entity.Account{}); tx.Error != nil {
		return tx.Error
	}
	return nil
}

func UpdateAccount(ids []int64, updates map[string]interface{}) error {
	if len(updates) > 0 {
		updates["update_time"] = time.Now()
		if tx := sqliteDB.Model(&entity.Account{}).
			Where("id in ?", ids).
			Updates(updates); tx.Error != nil {
			return tx.Error
		}
	}
	return nil
}

func GetAccount(query interface{}, args ...interface{}) (entity.Account, error) {
	var account entity.Account
	if tx := sqliteDB.Model(&entity.Account{}).
		Where(query, args...).First(&account); tx.Error != nil {
		if tx.Error == gorm.ErrRecordNotFound {
			return account, errors.New(constant.AccountNotExist)
		}
		return account, tx.Error
	}
	return account, nil
}

func PageAccount(accountPageDto dto.AccountPageDto) ([]entity.Account, int64, error) {
	var accounts []entity.Account
	var total int64
	tx := sqliteDB.Model(&entity.Account{}).
		Scopes(Paginate(accountPageDto.PageNum, accountPageDto.PageSize)).
		Order("create_time desc")
	if accountPageDto.Username != nil && *accountPageDto.Username != "" {
		tx.Where("username like ?", fmt.Sprintf("%%%s%%", *accountPageDto.Username))
	}
	if accountPageDto.Deleted != nil {
		tx.Where("deleted = ?", *accountPageDto.Deleted)
	}
	tx.Count(&total)
	if tx := tx.Find(&accounts); tx.Error != nil {
		return nil, 0, tx.Error
	}
	return accounts, total, nil
}

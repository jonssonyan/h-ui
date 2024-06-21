package dao

import (
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"h-ui/model/constant"
	"h-ui/model/dto"
	"h-ui/model/entity"
	"time"
)

func SaveAccount(account entity.Account) (int64, error) {
	if tx := sqliteDB.Save(&account); tx.Error != nil {
		logrus.Errorf("%v", tx.Error)
		return 0, errors.New(constant.SysError)
	}
	return *account.Id, nil
}

func DeleteAccount(ids []int64) error {
	if tx := sqliteDB.Where("id in ?", ids).Delete(&entity.Account{}); tx.Error != nil {
		logrus.Errorf("%v", tx.Error)
		return errors.New(constant.SysError)
	}
	return nil
}

func UpdateAccount(ids []int64, updates map[string]interface{}) error {
	if len(updates) > 0 {
		updates["update_time"] = time.Now().Format("2006-01-02 15:04:05")
		if tx := sqliteDB.Model(&entity.Account{}).
			Where("id in ?", ids).
			Updates(updates); tx.Error != nil {
			logrus.Errorf("%v", tx.Error)
			return errors.New(constant.SysError)
		}
	}
	return nil
}

func UpsertAccount(accounts []entity.Account) error {
	if tx := sqliteDB.Model(&entity.Account{}).Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "username"}},
		DoUpdates: clause.AssignmentColumns([]string{"pass", "con_pass", "quota", "download", "upload", "expire_time", "kick_util_time", "device_no", "role", "deleted", "create_time", "update_time"}),
	}).Create(accounts); tx.Error != nil {
		logrus.Errorf("%v", tx.Error)
		return errors.New(constant.SysError)
	}
	return nil
}

func UpdateAccountTraffic(username string, download int64, upload int64) error {
	if upload != 0 || download != 0 {
		updates := map[string]interface{}{}
		if download != 0 {
			updates["download"] = gorm.Expr("download + ?", download)
		}
		if upload != 0 {
			updates["upload"] = gorm.Expr("upload + ?", upload)
		}
		updates["update_time"] = time.Now().Format("2006-01-02 15:04:05")
		if tx := sqliteDB.Model(&entity.Account{}).
			Where("username = ?", username).
			Updates(updates); tx.Error != nil {
			logrus.Errorf("%v", tx.Error)
			return errors.New(constant.SysError)
		}
	}
	return nil
}

func GetAccount(query interface{}, args ...interface{}) (entity.Account, error) {
	var account entity.Account
	if tx := sqliteDB.Model(&entity.Account{}).
		Where(query, args...).First(&account); tx.Error != nil {
		if tx.Error == gorm.ErrRecordNotFound {
			return account, errors.New(constant.WrongPassword)
		}
		logrus.Errorf("%v", tx.Error)
		return account, errors.New(constant.SysError)
	}
	return account, nil
}

func PageAccount(accountPageDto dto.AccountPageDto) ([]entity.Account, int64, error) {
	var accounts []entity.Account
	var total int64
	tx := sqliteDB.Model(&entity.Account{}).
		Scopes(Paginate(accountPageDto.PageNum, accountPageDto.PageSize)).
		Order("role,create_time desc")
	if accountPageDto.Username != nil && *accountPageDto.Username != "" {
		tx.Where("username like ?", fmt.Sprintf("%%%s%%", *accountPageDto.Username))
	}
	if accountPageDto.Deleted != nil {
		tx.Where("deleted = ?", *accountPageDto.Deleted)
	}
	tx.Count(&total)
	if tx := tx.Find(&accounts); tx.Error != nil {
		logrus.Errorf("%v", tx.Error)
		return accounts, 0, errors.New(constant.SysError)
	}
	return accounts, total, nil
}

func ListAccount(query interface{}, args ...interface{}) ([]entity.Account, error) {
	var accounts []entity.Account
	if tx := sqliteDB.Model(&entity.Account{}).
		Where(query, args...).Order("role,create_time desc").Find(&accounts); tx.Error != nil {
		logrus.Errorf("%v", tx.Error)
		return accounts, errors.New(constant.SysError)
	}
	return accounts, nil
}

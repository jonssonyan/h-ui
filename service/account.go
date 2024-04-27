package service

import (
	"fmt"
	"h-ui/dao"
	"h-ui/model/bo"
	"h-ui/model/constant"
	"h-ui/model/dto"
	"h-ui/model/entity"
	"h-ui/util"
)

func Login(username string, pass string) (string, error) {
	account, err := dao.GetAccount("username = ? and pass = ?", username, util.SHA224String(pass))
	if err != nil {
		return "", err
	}
	accountBo := bo.AccountBo{
		Id:       *account.Id,
		Username: *account.Username,
		Deleted:  *account.Deleted,
	}
	return GenToken(accountBo)
}

func PageAccount(accountPageDto dto.AccountPageDto) ([]entity.Account, int64, error) {
	return dao.PageAccount(accountPageDto)
}

func SaveAccount(account entity.Account) error {
	_, err := dao.SaveAccount(account)
	return err
}

func DeleteAccount(ids []int64) error {
	return dao.DeleteAccount(ids)
}

func UpdateAccount(account entity.Account) error {
	var updates map[string]interface{}
	if account.Username != nil && *account.Username != "" {
		updates["username"] = *account.Username
	}
	if account.Pass != nil && *account.Pass != "" {
		updates["pass"] = util.SHA224String(*account.Pass)
	}
	if account.ConPass != nil && *account.ConPass != "" {
		updates["con_pass"] = util.SHA224String(fmt.Sprintf("%s@%s", *account.Username, *account.ConPass))
	}
	if account.Quota != nil {
		updates["quota"] = *account.Quota
	}
	if account.ExpireTime != nil {
		updates["expire_time"] = *account.ExpireTime
	}
	if account.Download != nil {
		updates["download"] = *account.Download
	}
	if account.Upload != nil {
		updates["upload"] = *account.Upload
	}
	if account.Deleted != nil {
		updates["deleted"] = *account.Deleted
	}
	return dao.UpdateAccount([]int64{*account.Id}, updates)
}

func ExistAccountUsername(username string) bool {
	_, err := dao.GetAccount("username = ?", username)
	if err != nil {
		if err.Error() == constant.AccountNotExist {
			return false
		}
	}
	return true
}

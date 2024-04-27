package service

import (
	"h-ui/dao"
	"h-ui/model/bo"
	"h-ui/model/dto"
	"h-ui/model/entity"
	"h-ui/util"
)

func PageAccount(accountPageDto dto.AccountPageDto) ([]entity.Account, int64, error) {
	return dao.PageAccount(accountPageDto)
}

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

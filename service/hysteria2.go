package service

import (
	"h-ui/dao"
)

func Hysteria2Auth(pass string) (string, error) {
	account, err := dao.GetAccount("deleted = 0 and pass = ? and CURRENT_TIMESTAMP < expire_time and (quota < 0 or quota > download + upload) ", pass)
	if err != nil {
		return "", err
	}
	return *account.Username, nil
}

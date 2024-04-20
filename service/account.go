package service

import (
	"h-ui/dao"
	"h-ui/model/dto"
	"h-ui/model/entity"
)

func PageAccount(accountPageDto dto.AccountPageDto) ([]entity.Account, int64, error) {
	return dao.PageAccount(accountPageDto)
}

package controller

import (
	"github.com/gin-gonic/gin"
	"h-ui/model/dto"
	"h-ui/model/vo"
	"h-ui/service"
)

func Login(c *gin.Context) {

}

func PageAccount(c *gin.Context) {
	accountPageDto, err := validateField(c, dto.AccountPageDto{})
	if err != nil {
		return
	}
	accounts, total, err := service.PageAccount(accountPageDto)
	if err != nil {
		vo.Fail(err.Error(), c)
		return
	}

	var accountVos []vo.AccountVo
	for _, item := range accounts {
		accountVo := vo.AccountVo{
			Username:   *item.Username,
			Quota:      *item.Quota,
			Download:   *item.Download,
			Upload:     *item.Upload,
			ExpireTime: *item.ExpireTime,
			Deleted:    *item.Deleted,
			BaseVo: vo.BaseVo{
				Id:         *item.Id,
				CreateTime: *item.CreateTime,
			},
		}
		accountVos = append(accountVos, accountVo)
	}
	accountPageVo := vo.AccountPageVo{
		AccountVos: accountVos,
		Total:      total,
	}
	vo.Success(accountPageVo, c)
}

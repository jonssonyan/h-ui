package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"h-ui/model/constant"
	"h-ui/model/dto"
	"h-ui/model/entity"
	"h-ui/model/vo"
	"h-ui/service"
	"h-ui/util"
	"time"
)

func Login(c *gin.Context) {
	loginDto, err := validateField(c, dto.LoginDto{})
	if err != nil {
		return
	}
	token, err := service.Login(*loginDto.Username, util.SHA224String(*loginDto.Pass))
	if err != nil {
		vo.Fail(err.Error(), c)
		return
	}
	jwtVo := vo.JwtVo{
		TokenType:   constant.TokenType,
		AccessToken: token,
	}
	vo.Success(jwtVo, c)
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
			Role:       *item.Role,
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

func SaveAccount(c *gin.Context) {
	accountSaveDto, err := validateField(c, dto.AccountSaveDto{})
	if err != nil {
		return
	}

	if service.ExistAccountUsername(*accountSaveDto.Username, 0) {
		vo.Fail(fmt.Sprintf("用户名%s已存在", *accountSaveDto.Username), c)
		return
	}

	passEncrypt := util.SHA224String(*accountSaveDto.Pass)
	conPassEncrypt := util.SHA224String(fmt.Sprintf("%s %s", *accountSaveDto.Username, *accountSaveDto.ConPass))
	account := entity.Account{
		Username:   accountSaveDto.Username,
		Pass:       &passEncrypt,
		ConPass:    &conPassEncrypt,
		Quota:      accountSaveDto.Quota,
		ExpireTime: accountSaveDto.ExpireTime,
		Deleted:    accountSaveDto.Deleted,
	}
	err = service.SaveAccount(account)
	if err != nil {
		vo.Fail(err.Error(), c)
		return
	}
	vo.Success(nil, c)
}

func DeleteAccount(c *gin.Context) {
	idDto, err := validateField(c, dto.IdDto{})
	if err != nil {
		return
	}
	account, err := service.GetAccount(*idDto.Id)
	if err != nil {
		vo.Fail(err.Error(), c)
		return
	}
	if *account.Role == "admin" {
		vo.Fail("admin 账号不能删除", c)
		return
	}
	err = service.DeleteAccount([]int64{*idDto.Id})
	if err != nil {
		vo.Fail(err.Error(), c)
		return
	}
	vo.Success(nil, c)
}

func UpdateAccount(c *gin.Context) {
	accountUpdateDto, err := validateField(c, dto.AccountUpdateDto{})
	if err != nil {
		return
	}

	if accountUpdateDto.Username != nil && *accountUpdateDto.Username != "" && service.ExistAccountUsername(*accountUpdateDto.Username, *accountUpdateDto.Id) {
		vo.Fail(fmt.Sprintf("用户名%s已存在", *accountUpdateDto.Username), c)
		return
	}

	var passEncrypt *string
	if accountUpdateDto.Pass != nil && *accountUpdateDto.Pass != "" {
		passEncryptSha224 := util.SHA224String(*accountUpdateDto.Pass)
		passEncrypt = &passEncryptSha224
	}

	var conPassEncrypt *string
	if accountUpdateDto.ConPass != nil && *accountUpdateDto.ConPass != "" {
		var username string
		if accountUpdateDto.Username != nil && *accountUpdateDto.Username != "" {
			username = *accountUpdateDto.Username
		} else {
			account, err := service.GetAccount(*accountUpdateDto.Id)
			if err != nil {
				vo.Fail(err.Error(), c)
				return
			}
			username = *account.Username
		}
		conPassSha224 := util.SHA224String(fmt.Sprintf("%s %s", username, *accountUpdateDto.ConPass))
		conPassEncrypt = &conPassSha224
	}

	account := entity.Account{
		Username:   accountUpdateDto.Username,
		Pass:       passEncrypt,
		ConPass:    conPassEncrypt,
		Quota:      accountUpdateDto.Quota,
		ExpireTime: accountUpdateDto.ExpireTime,
		Deleted:    accountUpdateDto.Deleted,
		BaseEntity: entity.BaseEntity{
			Id: accountUpdateDto.Id,
		},
	}
	if err = service.UpdateAccount(account); err != nil {
		vo.Fail(err.Error(), c)
		return
	}
	vo.Success(nil, c)
}

func GetAccountInfo(c *gin.Context) {
	myClaims, err := service.ParseToken(service.GetToken(c))
	if err != nil {
		vo.Fail(err.Error(), c)
		return
	}
	if myClaims.AccountBo.Deleted != 0 {
		vo.Fail("this account has been disabled", c)
		return
	}
	accountInfoVo := vo.AccountInfoVo{
		Id:       myClaims.AccountBo.Id,
		Username: myClaims.AccountBo.Username,
		Roles:    myClaims.AccountBo.Roles,
	}
	vo.Success(accountInfoVo, c)
}

func GetAccount(c *gin.Context) {
	idDto, err := validateField(c, dto.IdDto{})
	if err != nil {
		return
	}
	account, err := service.GetAccount(*idDto.Id)
	if err != nil {
		vo.Fail(err.Error(), c)
		return
	}
	accountVo := vo.AccountVo{
		BaseVo: vo.BaseVo{
			Id:         *account.Id,
			CreateTime: *account.CreateTime,
		},
		Username:   *account.Username,
		Quota:      *account.Quota,
		Download:   *account.Download,
		Upload:     *account.Upload,
		ExpireTime: *account.ExpireTime,
		Role:       *account.Role,
		Deleted:    *account.Deleted,
	}
	vo.Success(accountVo, c)
}

func ExportAccount(c *gin.Context) {
	accountExports, err := service.ListExportAccount()
	if err != nil {
		vo.Fail(err.Error(), c)
		return
	}

	fileName := fmt.Sprintf("AccountExport-%s.json", time.Now().Format("20060102150405"))
	filePath := constant.ExportPathDir + fileName

	if err = util.ExportJson(filePath, accountExports); err != nil {
		vo.Fail(err.Error(), c)
		return
	}

	// 下载
	if !util.Exists(filePath) {
		vo.Fail("file not exist", c)
		return
	}
	c.Header("Content-Type", "application/octet-stream")
	c.Header("Content-Transfer-Encoding", "binary")
	c.Header("Content-Disposition", fmt.Sprintf("attachment; filename=%s", fileName))
	c.File(filePath)
}

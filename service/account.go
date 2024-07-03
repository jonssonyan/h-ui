package service

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"h-ui/dao"
	"h-ui/model/bo"
	"h-ui/model/constant"
	"h-ui/model/dto"
	"h-ui/model/entity"
	"h-ui/model/vo"
)

func Login(username string, pass string) (string, error) {
	account, err := dao.GetAccount("username = ? and pass = ? and role = 'admin' and deleted = 0", username, pass)
	if err != nil {
		return "", err
	}
	accountBo := bo.AccountBo{
		Id:       *account.Id,
		Username: *account.Username,
		Roles:    []string{*account.Role},
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
	updates := map[string]interface{}{}
	if account.Username != nil && *account.Username != "" {
		updates["username"] = *account.Username
	}
	if account.Pass != nil && *account.Pass != "" {
		updates["pass"] = *account.Pass
	}
	if account.ConPass != nil && *account.ConPass != "" {
		updates["con_pass"] = fmt.Sprintf("%s.%s", *account.Username, *account.ConPass)
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
	if account.DeviceNo != nil {
		updates["device_no"] = *account.DeviceNo
	}
	if account.Deleted != nil {
		updates["deleted"] = *account.Deleted
	}
	return dao.UpdateAccount([]int64{*account.Id}, updates)
}

func ResetTraffic(id int64) error {
	return dao.UpdateAccount([]int64{id}, map[string]interface{}{"download": 0, "upload": 0})
}

func ExistAccountUsername(username string, id int64) bool {
	var err error
	if id != 0 {
		_, err = dao.GetAccount("username = ? and id != ?", username, id)
	} else {
		_, err = dao.GetAccount("username = ?", username)
	}
	if err != nil {
		if err.Error() == constant.WrongPassword {
			return false
		}
	}
	return true
}

func GetAccount(id int64) (entity.Account, error) {
	return dao.GetAccount("id = ?", id)
}

func ListExportAccount() ([]bo.AccountExport, error) {
	accounts, err := dao.ListAccount(nil, nil)
	if err != nil {
		return nil, errors.New(constant.SysError)
	}
	var accountExports []bo.AccountExport
	for _, item := range accounts {
		accountExport := bo.AccountExport{
			Id:           *item.Id,
			Username:     *item.Username,
			Pass:         *item.Pass,
			ConPass:      *item.ConPass,
			Quota:        *item.Quota,
			Download:     *item.Download,
			Upload:       *item.Upload,
			ExpireTime:   *item.ExpireTime,
			DeviceNo:     *item.DeviceNo,
			KickUtilTime: *item.KickUtilTime,
			Role:         *item.Role,
			Deleted:      *item.Deleted,
			CreateTime:   *item.CreateTime,
			UpdateTime:   *item.UpdateTime,
		}
		accountExports = append(accountExports, accountExport)
	}
	return accountExports, nil
}

func ReleaseKickAccount(id int64) error {
	return dao.UpdateAccount([]int64{id}, map[string]interface{}{"kick_util_time": 0})
}

func UpsertAccount(accounts []entity.Account) error {
	return dao.UpsertAccount(accounts)
}

func GetAccountInfo(c *gin.Context) (vo.AccountInfoVo, error) {
	myClaims, err := ParseToken(GetToken(c))
	if err != nil {
		return vo.AccountInfoVo{}, err
	}
	if myClaims.AccountBo.Deleted != 0 {
		return vo.AccountInfoVo{}, errors.New("this account has been disabled")
	}
	return vo.AccountInfoVo{
		Id:       myClaims.AccountBo.Id,
		Username: myClaims.AccountBo.Username,
		Roles:    myClaims.AccountBo.Roles,
	}, nil
}

package service

import (
	"h-ui/dao"
	"h-ui/model/constant"
	"h-ui/util"
)

func InitPortHopping() error {
	if err := util.RemoveByComment(); err != nil {
		return err
	}

	hysteria2Config, err := GetHysteria2Config()
	if err != nil {
		return err
	}

	// set port forward
	hysteria2ConfigPortHopping, err := dao.GetConfig("key = ?", constant.Hysteria2ConfigPortHopping)
	if err != nil {
		return err
	}
	if *hysteria2ConfigPortHopping.Value != "" {
		if err := util.PortForward(*hysteria2ConfigPortHopping.Value, *hysteria2Config.Listen, util.Add); err != nil {
			return err
		}
	}
	return nil
}

package service

import (
	"h-ui/dao"
	"h-ui/model/constant"
	"h-ui/util"
	"strings"
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
		listen := strings.Split(*hysteria2Config.Listen, ":")
		if len(listen) == 2 {
			if err := util.PortForward(*hysteria2ConfigPortHopping.Value, listen[1], util.Add); err != nil {
				return err
			}
		}
	}
	return nil
}

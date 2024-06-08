package controller

import (
	"github.com/gin-gonic/gin"
	"h-ui/model/dto"
	"h-ui/model/vo"
	"h-ui/service"
	"h-ui/util"
	"os"
	"strings"
)

func Hysteria2Auth(c *gin.Context) {
	hysteria2AuthDto, err := validateField(c, dto.Hysteria2AuthDto{})
	if err != nil {
		return
	}
	username, err := service.Hysteria2Auth(*hysteria2AuthDto.Auth)
	if err != nil || username == "" {
		vo.Hysteria2AuthFail("", c)
		return
	}
	vo.Hysteria2AuthSuccess(username, c)
}

func Hysteria2Kick(c *gin.Context) {
	hysteria2KickDto, err := validateField(c, dto.Hysteria2KickDto{})
	if err != nil {
		return
	}
	err = service.Hysteria2Kick(hysteria2KickDto.Ids, *hysteria2KickDto.KickUtilTime)
	if err != nil {
		vo.Fail(err.Error(), c)
		return
	}
	vo.Success(nil, c)
}

func ChangeHysteria2Version(c *gin.Context) {
	hysteria2VersionDto, err := validateField(c, dto.Hysteria2VersionDto{})
	if err != nil {
		return
	}
	if err = service.StopHysteria2(); err != nil {
		vo.Fail(err.Error(), c)
		return
	}

	if util.Exists(util.GetHysteria2BinPath()) {
		if err := os.Remove(util.GetHysteria2BinPath()); err != nil {
			vo.Fail("filed remove hysteria2 bin file", c)
			return
		}
	}

	if err = util.DownloadHysteria2(*hysteria2VersionDto.Version); err != nil {
		vo.Fail(err.Error(), c)
		return
	}

	if err = service.StartHysteria2(); err != nil {
		vo.Fail(err.Error(), c)
		return
	}
	vo.Success(nil, c)
}

func ListRelease(c *gin.Context) {
	releases, err := util.ListRelease("apernet", "hysteria")
	if err != nil {
		vo.Fail(err.Error(), c)
		return
	}

	var vos []vo.Hysteria2ReleaseVo
	for _, item := range releases {
		versionSplit := strings.Split(*item.TagName, "/v")
		if len(versionSplit) == 2 {
			// >= v2.4.4
			result := util.CompareVersion(versionSplit[1], "2.4.4")
			if result < 0 {
				break
			}
			for _, asset := range item.Assets {
				if asset.GetName() == util.GetHysteria2BinName() {
					vos = append(vos, vo.Hysteria2ReleaseVo{
						TagName:            *item.TagName,
						BrowserDownloadURL: asset.GetBrowserDownloadURL(),
					})
					break
				}
			}
		}
	}
	vo.Success(vos, c)
}

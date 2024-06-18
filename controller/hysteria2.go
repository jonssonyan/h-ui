package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"h-ui/model/constant"
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

func Hysteria2Subscribe(c *gin.Context) {
	conPass := c.Param("conPass")
	userAgent := c.Request.Header.Get("User-Agent")
	host := c.Request.Header.Get("Host")

	if host == "" {
		vo.Fail("Host is empty", c)
		return
	}

	var clientType string
	if strings.HasPrefix(userAgent, "ClashforWindows") {
		clientType = "ClashforWindows"
	} else if strings.HasPrefix(userAgent, "Shadowrocket") {
		clientType = "Shadowrocket"
	} else if strings.HasPrefix(userAgent, "v2rayN") {
		clientType = "v2rayN"
	}

	userInfo, configYaml, err := service.Hysteria2Subscribe(conPass, clientType, host)
	if err != nil {
		vo.Fail(err.Error(), c)
		return
	}

	if clientType == "ClashforWindows" || clientType == "Shadowrocket" {
		c.Header("content-disposition", "attachment; filename=hui.yaml")
		c.Header("profile-update-interval", "12")
		c.Header("subscription-userinfo", userInfo)
	}

	c.String(200, configYaml)
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

func Hysteria2ChangeVersion(c *gin.Context) {
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

	if err = util.DownloadHysteria2(fmt.Sprintf("app/%s", *hysteria2VersionDto.Version)); err != nil {
		vo.Fail(err.Error(), c)
		return
	}

	config, err := service.GetConfig(constant.Hysteria2Enable)
	if err != nil {
		vo.Fail(err.Error(), c)
		return
	}

	if *config.Value == "1" {
		if err = service.StartHysteria2(); err != nil {
			vo.Fail(err.Error(), c)
			return
		}
	}
	vo.Success(nil, c)
}

func ListRelease(c *gin.Context) {
	releases, err := util.ListRelease("apernet", "hysteria")
	if err != nil {
		vo.Fail("get Hysteria2 releases failed", c)
		return
	}

	var versions []string
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
					versions = append(versions, fmt.Sprintf("v%s", versionSplit[1]))
					break
				}
			}
		}
	}
	vo.Success(versions, c)
}

func Hysteria2SubscribeUrl(c *gin.Context) {
	hysteria2SubscribeUrlDto, err := validateField(c, dto.Hysteria2SubscribeUrlDto{})
	if err != nil {
		return
	}
	subscribeUrl, err := service.Hysteria2SubscribeUrl(*hysteria2SubscribeUrlDto.AccountId,
		*hysteria2SubscribeUrlDto.Protocol,
		*hysteria2SubscribeUrlDto.Host)
	if err != nil {
		vo.Fail(err.Error(), c)
		return
	}
	vo.Success(subscribeUrl, c)
}

func Hysteria2Url(c *gin.Context) {
	hysteria2UrlDto, err := validateField(c, dto.Hysteria2UrlDto{})
	if err != nil {
		return
	}
	url, err := service.Hysteria2Url(*hysteria2UrlDto.AccountId, *hysteria2UrlDto.Hostname)
	if err != nil {
		vo.Fail(err.Error(), c)
		return
	}
	vo.Success(url, c)
}

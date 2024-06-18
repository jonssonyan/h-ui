package controller

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
	"h-ui/dao"
	"h-ui/model/bo"
	"h-ui/model/constant"
	"h-ui/model/dto"
	"h-ui/model/entity"
	"h-ui/model/vo"
	"h-ui/service"
	"h-ui/util"
	"io"
	"strings"
	"time"
)

func UpdateConfigs(c *gin.Context) {
	configsUpdateDto, err := validateField(c, dto.ConfigsUpdateDto{})
	if err != nil {
		return
	}

	config, err := service.GetConfig(constant.HUIWebPort)
	if err != nil {
		vo.Fail(err.Error(), c)
		return
	}

	for _, item := range configsUpdateDto.ConfigUpdateDtos {
		if err = service.UpdateConfig(*item.Key, *item.Value); err != nil {
			vo.Fail(err.Error(), c)
			return
		}
	}

	for _, item := range configsUpdateDto.ConfigUpdateDtos {
		if item.Key != nil && *item.Key == constant.HUIWebPort && *config.Value != *item.Value {
			go func() {
				if err := service.StartServer(gin.Default()); err != nil {
					logrus.Errorf(err.Error())
					// 启动失败后将端口改为以前的端口
					if err := service.UpdateConfig(constant.HUIWebPort, *config.Value); err != nil {
						logrus.Errorf(err.Error())
					}
				}
			}()
			break
		}
	}
	vo.Success(nil, c)
}

func GetConfig(c *gin.Context) {
	configDto, err := validateField(c, dto.ConfigDto{})
	if err != nil {
		return
	}
	config, err := service.GetConfig(*configDto.Key)
	if err != nil {
		vo.Fail(err.Error(), c)
		return
	}
	configVo := vo.ConfigVo{
		Key:   *config.Key,
		Value: *config.Value,
	}

	running := service.Hysteria2Status()

	if (*config.Value == "1") != running {
		enable := "0"
		if running {
			enable = "1"
		}
		if err := service.UpdateConfig(constant.Hysteria2Enable, enable); err != nil {
			vo.Fail(err.Error(), c)
			return
		}
		configVo.Value = enable
	}

	vo.Success(configVo, c)
}

func ListConfig(c *gin.Context) {
	configsDto, err := validateField(c, dto.ConfigsDto{})
	if err != nil {
		return
	}
	configs, err := service.ListConfig(configsDto.Keys)
	if err != nil {
		vo.Fail(err.Error(), c)
		return
	}
	var configVos []vo.ConfigVo
	for _, item := range configs {
		configVo := vo.ConfigVo{
			Key:   *item.Key,
			Value: *item.Value,
		}
		configVos = append(configVos, configVo)
	}
	vo.Success(configVos, c)
}

func GetHysteria2Config(c *gin.Context) {
	config, err := service.GetHysteria2Config()
	if err != nil {
		vo.Fail(err.Error(), c)
		return
	}
	vo.Success(config, c)
}

func UpdateHysteria2Config(c *gin.Context) {
	hysteria2ServerConfig, err := validateField(c, bo.Hysteria2ServerConfig{})
	if err != nil {
		return
	}
	if err = service.UpdateHysteria2Config(hysteria2ServerConfig); err != nil {
		vo.Fail(err.Error(), c)
		return
	}
	vo.Success(nil, c)
}

func ExportHysteria2Config(c *gin.Context) {
	hysteria2ServerConfig, err := service.GetHysteria2Config()
	if err != nil {
		vo.Fail(err.Error(), c)
		return
	}

	// 默认值
	getConfig, err := dao.GetConfig("key = ?", constant.HUIWebPort)
	if err != nil {
		vo.Fail(err.Error(), c)
		return
	}
	authType := "http"
	authHttpUrl := fmt.Sprintf("http://127.0.0.1:%s/hui/hysteria2/auth", *getConfig.Value)
	authHttpInsecure := true
	trafficStatsSecret := ""
	var auth bo.ServerConfigAuth
	auth.Type = &authType
	var http bo.ServerConfigAuthHTTP
	http.URL = &authHttpUrl
	http.Insecure = &authHttpInsecure
	auth.HTTP = &http
	var trafficStats bo.ServerConfigTrafficStats
	trafficStats.Secret = &trafficStatsSecret
	hysteria2ServerConfig.Auth = &auth
	hysteria2ServerConfig.TrafficStats = &trafficStats

	fileName := fmt.Sprintf("Hysteria2Config-%s.yaml", time.Now().Format("20060102150405"))
	filePath := constant.ExportPathDir + fileName

	if err = util.ExportFile(filePath, hysteria2ServerConfig, 1); err != nil {
		vo.Fail(err.Error(), c)
		return
	}

	if !util.Exists(filePath) {
		vo.Fail("file not exist", c)
		return
	}
	c.Header("Content-Type", "application/octet-stream")
	c.Header("Content-Transfer-Encoding", "binary")
	c.Header("Content-Disposition", fmt.Sprintf("attachment; filename=%s", fileName))
	c.File(filePath)
}

func ImportHysteria2Config(c *gin.Context) {
	file, header, err := c.Request.FormFile("file")
	if err != nil {
		vo.Fail(constant.SysError, c)
		return
	}
	if header.Size > 1024*1024*2 {
		vo.Fail("the file is too big", c)
		return
	}
	if !strings.HasSuffix(header.Filename, ".yaml") {
		vo.Fail("file format not supported", c)
		return
	}
	content, err := io.ReadAll(file)
	if err != nil {
		vo.Fail("yaml file read err", c)
		return
	}
	var hysteria2ServerConfig bo.Hysteria2ServerConfig
	if err = yaml.Unmarshal(content, &hysteria2ServerConfig); err != nil {
		vo.Fail("content Unmarshal err", c)
		return
	}

	// 默认值
	getConfig, err := dao.GetConfig("key = ?", constant.HUIWebPort)
	if err != nil {
		vo.Fail(err.Error(), c)
		return
	}
	authType := "http"
	authHttpUrl := fmt.Sprintf("http://127.0.0.1:%s/hui/hysteria2/auth", *getConfig.Value)
	authHttpInsecure := true
	trafficStatsSecret := ""
	var auth bo.ServerConfigAuth
	auth.Type = &authType
	var http bo.ServerConfigAuthHTTP
	http.URL = &authHttpUrl
	http.Insecure = &authHttpInsecure
	auth.HTTP = &http
	var trafficStats bo.ServerConfigTrafficStats
	trafficStats.Secret = &trafficStatsSecret
	hysteria2ServerConfig.Auth = &auth
	hysteria2ServerConfig.TrafficStats = &trafficStats

	if err = service.SetHysteria2Config(hysteria2ServerConfig); err != nil {
		vo.Fail(err.Error(), c)
		return
	}
	vo.Success(nil, c)
}

func ExportConfig(c *gin.Context) {
	configs, err := service.ListConfigNotIn([]string{constant.Hysteria2Config})
	if err != nil {
		vo.Fail(err.Error(), c)
		return
	}
	fileName := fmt.Sprintf("SystemConfig-%s.json", time.Now().Format("20060102150405"))
	filePath := constant.ExportPathDir + fileName

	if err = util.ExportFile(filePath, configs, 0); err != nil {
		vo.Fail(err.Error(), c)
		return
	}

	if !util.Exists(filePath) {
		vo.Fail("file not exist", c)
		return
	}
	c.Header("Content-Type", "application/octet-stream")
	c.Header("Content-Transfer-Encoding", "binary")
	c.Header("Content-Disposition", fmt.Sprintf("attachment; filename=%s", fileName))
	c.File(filePath)
}

func ImportConfig(c *gin.Context) {
	file, header, err := c.Request.FormFile("file")
	if err != nil {
		vo.Fail(constant.SysError, c)
		return
	}
	if header.Size > 1024*1024*2 {
		vo.Fail("the file is too big", c)
		return
	}
	if !strings.HasSuffix(header.Filename, ".json") {
		vo.Fail("file format not supported", c)
		return
	}
	content, err := io.ReadAll(file)
	if err != nil {
		vo.Fail("json file read err", c)
		return
	}
	var configs []entity.Config
	if err = json.Unmarshal(content, &configs); err != nil {
		vo.Fail("content Unmarshal err", c)
		return
	}
	if err = service.UpsertConfig(configs); err != nil {
		vo.Fail(err.Error(), c)
		return
	}
	vo.Success(nil, c)
}

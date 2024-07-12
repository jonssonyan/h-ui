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
	"regexp"
	"strconv"
	"strings"
	"time"
)

func UpdateConfigs(c *gin.Context) {
	configsUpdateDto, err := validateField(c, dto.ConfigsUpdateDto{})
	if err != nil {
		return
	}

	port, crtPath, keyPath, err := service.GetPortAndCert()
	if err != nil {
		vo.Fail(err.Error(), c)
		return
	}

	needResetPortHopping := false
	needRestart := false

	for _, item := range configsUpdateDto.ConfigUpdateDtos {
		key := *item.Key
		value := *item.Value

		if key == constant.HUIWebPort && strconv.FormatInt(port, 10) != value {
			port, err := strconv.Atoi(value)
			if err != nil {
				vo.Fail(fmt.Sprintf("port: %s is invalid", value), c)
				return
			}
			if !util.IsPortAvailable(uint(port), "tcp") {
				vo.Fail(fmt.Sprintf("port: %s is used", value), c)
				return
			}
			needRestart = true
		}
		if key == constant.HUICrtPath && crtPath != value {
			if value != "" && !util.Exists(value) {
				vo.Fail(fmt.Sprintf("crt path: %s is not exist", value), c)
				return
			}
			needRestart = true
		}
		if key == constant.HUIKeyPath && keyPath != value {
			if value != "" && !util.Exists(value) {
				vo.Fail(fmt.Sprintf("key path: %s is not exist", value), c)
				return
			}
			needRestart = true
		}

		if key == constant.Hysteria2ConfigPortHopping {
			re := regexp.MustCompile("^\\d+(?:-\\d+)?(?:,\\d+(?:-\\d+)?)*$")
			if value != "" && !re.MatchString(value) {
				vo.Fail(fmt.Sprintf("port hopping: %s is invalid", value), c)
				return
			}
			hysteria2ConfigPortHopping, err := service.GetConfig(constant.Hysteria2ConfigPortHopping)
			if err != nil {
				vo.Fail(err.Error(), c)
				return
			}
			if *hysteria2ConfigPortHopping.Value != value {
				needResetPortHopping = true
			}
		}
		if err = service.UpdateConfig(key, value); err != nil {
			vo.Fail(err.Error(), c)
			return
		}
	}

	if needResetPortHopping {
		if err := service.InitPortHopping(); err != nil {
			vo.Fail(err.Error(), c)
			return
		}
	}

	if needRestart {
		go func() {
			_ = service.StopServer()
		}()
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

	running := service.Hysteria2IsRunning()

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

	hysteria2Config, err := service.GetHysteria2Config()
	if err != nil {
		vo.Fail(err.Error(), c)
		return
	}

	needResetPortHopping := false
	if hysteria2Config.Listen != nil &&
		*hysteria2Config.Listen != "" &&
		hysteria2ServerConfig.Listen != nil &&
		*hysteria2ServerConfig.Listen != "" &&
		*hysteria2ServerConfig.Listen != *hysteria2Config.Listen {
		needResetPortHopping = true
	}

	if err = service.UpdateHysteria2Config(hysteria2ServerConfig); err != nil {
		vo.Fail(err.Error(), c)
		return
	}

	if needResetPortHopping {
		if err := service.InitPortHopping(); err != nil {
			vo.Fail(err.Error(), c)
			return
		}
	}

	running := service.Hysteria2IsRunning()
	if running {
		if err = service.RestartHysteria2(); err != nil {
			vo.Fail(err.Error(), c)
			return
		}
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
	config, err := dao.ListConfig("key in ?", []string{constant.HUIWebPort, constant.JwtSecret})
	if err != nil {
		vo.Fail(err.Error(), c)
		return
	}

	var hUIWebPort string
	var jwtSecret string
	for _, item := range config {
		if *item.Key == constant.HUIWebPort {
			hUIWebPort = *item.Value
		} else if *item.Key == constant.JwtSecret {
			jwtSecret = *item.Value
		}
	}

	if hUIWebPort == "" || jwtSecret == "" {
		logrus.Errorf("hUIWebPort or jwtSecret is nil")
		vo.Fail(constant.SysError, c)
		return
	}

	authHttpUrl, err := service.GetAuthHttpUrl()
	if err != nil {
		vo.Fail(err.Error(), c)
		return
	}

	authType := "http"
	authHttpInsecure := true
	var auth bo.ServerConfigAuth
	auth.Type = &authType
	var http bo.ServerConfigAuthHTTP
	http.URL = &authHttpUrl
	http.Insecure = &authHttpInsecure
	auth.HTTP = &http
	hysteria2ServerConfig.Auth = &auth
	hysteria2ServerConfig.TrafficStats.Secret = &jwtSecret

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
	config, err := dao.ListConfig("key in ?", []string{constant.HUIWebPort, constant.JwtSecret})
	if err != nil {
		vo.Fail(err.Error(), c)
		return
	}

	var hUIWebPort string
	var jwtSecret string
	for _, item := range config {
		if *item.Key == constant.HUIWebPort {
			hUIWebPort = *item.Value
		} else if *item.Key == constant.JwtSecret {
			jwtSecret = *item.Value
		}
	}

	if hUIWebPort == "" || jwtSecret == "" {
		logrus.Errorf("hUIWebPort or jwtSecret is nil")
		vo.Fail(constant.SysError, c)
		return
	}

	authHttpUrl, err := service.GetAuthHttpUrl()
	if err != nil {
		vo.Fail(err.Error(), c)
		return
	}

	authType := "http"
	authHttpInsecure := true
	var auth bo.ServerConfigAuth
	auth.Type = &authType
	var http bo.ServerConfigAuthHTTP
	http.URL = &authHttpUrl
	http.Insecure = &authHttpInsecure
	auth.HTTP = &http
	hysteria2ServerConfig.Auth = &auth
	hysteria2ServerConfig.TrafficStats.Secret = &jwtSecret

	if err = service.SetHysteria2Config(hysteria2ServerConfig); err != nil {
		vo.Fail(err.Error(), c)
		return
	}

	running := service.Hysteria2IsRunning()
	if running {
		if err = service.RestartHysteria2(); err != nil {
			vo.Fail(err.Error(), c)
			return
		}
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

func Hysteria2AcmePath(c *gin.Context) {
	hysteria2AcmePathVo, err := service.Hysteria2AcmePath()
	if err != nil {
		vo.Fail(err.Error(), c)
		return
	}
	vo.Success(hysteria2AcmePathVo, c)
}

func RestartServer(c *gin.Context) {
	go func() {
		_ = service.StopServer()
	}()
	vo.Success(nil, c)
}

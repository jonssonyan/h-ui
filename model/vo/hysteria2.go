package vo

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type hysteria2Result struct {
	Ok bool   `json:"ok"`
	Id string `json:"id"`
}

func Hysteria2AuthSuccess(id string, c *gin.Context) {
	c.JSON(http.StatusOK, hysteria2Result{
		Ok: true,
		Id: id,
	})
}

func Hysteria2AuthFail(id string, c *gin.Context) {
	c.JSON(http.StatusOK, hysteria2Result{
		Ok: false,
		Id: id,
	})
}

type Hysteria2SubscribeVo struct {
	Url    string `json:"url"`
	QrCode []byte `json:"qrCode"`
}

type Hysteria2UrlVo struct {
	Url    string `json:"url"`
	QrCode []byte `json:"qrCode"`
}

type Hysteria2AcmePathVo struct {
	CrtPath string `json:"crtPath"`
	KeyPath string `json:"keyPath"`
}

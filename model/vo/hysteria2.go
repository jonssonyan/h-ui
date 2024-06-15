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

type Hysteria2ReleaseVo struct {
	TagName            string `json:"tagName"`
	BrowserDownloadURL string `json:"browserDownloadURL"`
}

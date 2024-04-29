package frontend

import (
	"embed"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/fs"
	"net/http"
)

//go:embed dist/*
var staticFiles embed.FS

func InitFrontend(router *gin.Engine) {

	router.GET("/", func(c *gin.Context) {
		indexHTML, err := staticFiles.ReadFile("dist/index.html")
		if err != nil {
			c.String(http.StatusInternalServerError, "Internal Server Error")
			return
		}
		c.Data(http.StatusOK, "text/html", indexHTML)
	})

	router.StaticFS("/static", http.FS(getStaticFS()))

	router.NoRoute(func(c *gin.Context) {
		filePath := c.Request.URL.Path
		fileContent, err := getFileContent(filePath)
		if err != nil {
			c.String(http.StatusNotFound, "404")
			return
		}
		c.Data(http.StatusOK, http.DetectContentType(fileContent), fileContent)
	})
}

func getStaticFS() fs.FS {
	staticFs, _ := fs.Sub(staticFiles, "dist/static")
	return staticFs
}

func getFileContent(filePath string) ([]byte, error) {
	fileContent, err := staticFiles.ReadFile(fmt.Sprintf("dist%s", filePath))
	if err != nil {
		return nil, err
	}
	return fileContent, nil
}

package frontend

import (
	"embed"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/fs"
	"net/http"
	"path"
	"strings"
)

//go:embed dist/*
var staticFiles embed.FS

func InitFrontend(router *gin.Engine, relativePath string) {
	router.GET(relativePath, func(c *gin.Context) {
		indexHTML, err := staticFiles.ReadFile("dist/index.html")
		if err != nil {
			c.String(http.StatusInternalServerError, "Internal Server Error")
			return
		}
		c.Data(http.StatusOK, "text/html", indexHTML)
	})

	router.GET(path.Join(relativePath, "favicon.ico"), func(c *gin.Context) {
		indexHTML, err := staticFiles.ReadFile("dist/favicon.ico")
		if err != nil {
			c.String(http.StatusInternalServerError, "Internal Server Error")
			return
		}
		c.Data(http.StatusOK, "image/x-icon", indexHTML)
	})

	router.StaticFS(path.Join(relativePath, "assets"), http.FS(getStaticFS()))

	router.NoRoute(func(c *gin.Context) {
		filePath := c.Request.URL.Path
		if relativePath != "/" && !strings.HasPrefix(filePath, relativePath) {
			c.String(http.StatusNotFound, "404")
		}
		fileContent, err := getFileContent(filePath)
		if err != nil {
			c.String(http.StatusNotFound, "404")
			return
		}
		c.Data(http.StatusOK, http.DetectContentType(fileContent), fileContent)
	})
}

func getStaticFS() fs.FS {
	staticFs, _ := fs.Sub(staticFiles, "dist/assets")
	return staticFs
}

func getFileContent(filePath string) ([]byte, error) {
	fileContent, err := staticFiles.ReadFile(fmt.Sprintf("dist%s", filePath))
	if err != nil {
		return nil, err
	}
	return fileContent, nil
}

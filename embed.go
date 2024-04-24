package main

import (
	"embed"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/fs"
	"net/http"
)

//go:embed frontend/build/*
var staticFiles embed.FS

func InitFrontend(router *gin.Engine) {

	router.GET("/", func(c *gin.Context) {
		indexHTML, err := staticFiles.ReadFile("frontend/build/index.html")
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
	staticFs, _ := fs.Sub(staticFiles, "frontend/build/static")
	return staticFs
}

func getFileContent(filePath string) ([]byte, error) {
	fileContent, err := staticFiles.ReadFile(fmt.Sprintf("frontend/build%s", filePath))
	if err != nil {
		return nil, err
	}
	return fileContent, nil
}

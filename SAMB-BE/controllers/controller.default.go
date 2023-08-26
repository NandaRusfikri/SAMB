package controllers

import (
	"SAMB-BE/constant"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type DefaultController struct{}

func InitControllerDefault(g *gin.Engine) {
	handler := &DefaultController{}

	g.GET("/", handler.MainPage)
	g.Static("/assets/", "./assets/")
	g.Static("/storage", "./storage")
}

func (h *DefaultController) MainPage(c *gin.Context) {
	jsonData := map[string]interface{}{
		"application_name": constant.SERVICE_NAME,
		"version":          constant.SERVICE_VERSION,
		"time_server":      time.Now(),
	}

	c.JSON(http.StatusOK, jsonData)
}

package web

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/aceberg/WatchYourPorts/internal/models"
)

func historyHandler(c *gin.Context) {
	var guiData models.GuiData
	guiData.Config = appConfig

	c.HTML(http.StatusOK, "header.html", guiData)
	c.HTML(http.StatusOK, "history.html", guiData)
}

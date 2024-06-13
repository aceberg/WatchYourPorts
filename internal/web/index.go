package web

import (
	"net/http"

	"github.com/gin-gonic/gin"

	// "github.com/aceberg/WatchYourPorts/internal/db"
	"github.com/aceberg/WatchYourPorts/internal/models"
)

func indexHandler(c *gin.Context) {
	var guiData models.GuiData
	guiData.Config = appConfig

	guiData.AddrMap = allAddrs

	c.HTML(http.StatusOK, "header.html", guiData)
	c.HTML(http.StatusOK, "index.html", guiData)
}

package web

import (
	// "log"
	"net/http"

	"github.com/gin-gonic/gin"

	// "github.com/aceberg/WatchYourPorts/internal/db"
	"github.com/aceberg/WatchYourPorts/internal/models"
)

func indexHandler(c *gin.Context) {
	var guiData models.GuiData
	guiData.Config = appConfig

	for key, value := range allAddrs {
		value.Total = 0
		value.Watching = 0
		value.Online = 0
		value.Offline = 0

		for _, p := range value.PortMap {
			value.Total++
			if p.Watch {
				value.Watching++
			}
			if p.State {
				value.Online++
			} else {
				value.Offline++
			}
		}
		// log.Println("KEY =", key)
		// log.Println("VALUE =", value)

		allAddrs[key] = value
	}

	guiData.AddrMap = allAddrs

	c.HTML(http.StatusOK, "header.html", guiData)
	c.HTML(http.StatusOK, "index.html", guiData)
}

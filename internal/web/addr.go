package web

import (
	// "log"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/aceberg/WatchYourPorts/internal/models"
	"github.com/aceberg/WatchYourPorts/internal/yaml"
)

func addHandler(c *gin.Context) {

	addr := c.PostForm("addr")
	name := c.PostForm("name")

	if addr != "" {

		oneAddr := models.AddrToScan{}
		oneAddr.Addr = addr
		oneAddr.Name = name

		allAddrs[addr] = oneAddr

		yaml.Write(appConfig.YamlPath, allAddrs)
	}

	c.Redirect(http.StatusFound, "/")
}

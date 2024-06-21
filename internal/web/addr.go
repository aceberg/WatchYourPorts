package web

import (
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

func renameHandler(c *gin.Context) {

	addr := c.PostForm("addr")
	name := c.PostForm("name")

	if addr != "" {

		oneAddr := allAddrs[addr]
		oneAddr.Name = name

		allAddrs[addr] = oneAddr

		yaml.Write(appConfig.YamlPath, allAddrs)
	}

	c.Redirect(http.StatusFound, "/scan/?addr="+addr)
}

func delHandler(c *gin.Context) {

	addr := c.PostForm("addr")

	if addr != "" {

		delete(allAddrs, addr)

		yaml.Write(appConfig.YamlPath, allAddrs)
	}

	c.Redirect(http.StatusFound, "/")
}

package web

import (
	"net/http"
	"slices"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/aceberg/WatchYourPorts/internal/models"
	"github.com/aceberg/WatchYourPorts/internal/yaml"
)

func scanHandler(c *gin.Context) {
	var guiData models.GuiData
	guiData.Config = appConfig

	addr, ok := c.GetQuery("addr")
	if ok {
		guiData.OneAddr = allAddrs[addr]
	}

	c.HTML(http.StatusOK, "header.html", guiData)
	c.HTML(http.StatusOK, "scan.html", guiData)
}

func scanSaveHandler(c *gin.Context) {

	addr := c.PostForm("addr")

	names := c.PostFormArray("name")
	ports := c.PostFormArray("port")
	states := c.PostFormArray("state")
	watchs := c.PostFormArray("watch")

	tmpMap := make(map[int]models.PortItem)
	onePort := models.PortItem{}

	for i, port := range ports {
		if slices.Contains(watchs, port) {
			onePort.Watch = true
		} else {
			onePort.Watch = false
		}

		onePort.Name = names[i]
		onePort.Port, _ = strconv.Atoi(port)
		onePort.State, _ = strconv.ParseBool(states[i])

		tmpMap[onePort.Port] = onePort
	}

	oneAddr := allAddrs[addr]
	oneAddr.PortMap = tmpMap
	allAddrs[addr] = oneAddr

	// log.Println("SAVEALL:", allAddrs)

	yaml.Write(appConfig.YamlPath, allAddrs)

	c.Redirect(http.StatusFound, "/scan/?addr="+addr)
}

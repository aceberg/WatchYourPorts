package web

import (
	// "log"
	"net/http"
	"slices"
	"strconv"

	"github.com/gin-gonic/gin"

	// "github.com/aceberg/WatchYourPorts/internal/db"
	"github.com/aceberg/WatchYourPorts/internal/models"
	"github.com/aceberg/WatchYourPorts/internal/port"
	"github.com/aceberg/WatchYourPorts/internal/yaml"
)

func scanHandler(c *gin.Context) {
	var guiData models.GuiData
	guiData.Config = appConfig

	addr, ok := c.GetQuery("addr")
	if ok {
		for _, oneAddr := range allAddrs {
			if oneAddr.Addr == addr {
				guiData.OneAddr = oneAddr
				break
			}
		}

	}

	c.HTML(http.StatusOK, "header.html", guiData)
	c.HTML(http.StatusOK, "scan.html", guiData)
}

func scanPortsHandler(c *gin.Context) {
	addr := c.PostForm("addr")
	begin := c.PostForm("begin")
	end := c.PostForm("end")

	if begin == "" {
		begin = "1"
	}
	if end == "" {
		end = "65535"
	}
	ports := port.Scan(addr, begin, end)
	// log.Println("PORTS:", ports)

	tmpMap := allAddrs[addr].PortMap

	var exists bool
	var onePort models.PortItem

	for _, p := range ports {
		onePort, exists = tmpMap[p]
		if !exists {
			onePort.Port = p
			onePort.State = true
			tmpMap[p] = onePort
		}
	}

	oneAddr := allAddrs[addr]
	oneAddr.PortMap = tmpMap
	allAddrs[addr] = oneAddr

	c.Redirect(http.StatusFound, "/scan/?addr="+addr)
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

			onePort.Name = names[i]
			onePort.Port, _ = strconv.Atoi(port)
			onePort.State, _ = strconv.ParseBool(states[i])
			onePort.Watch = true

			tmpMap[onePort.Port] = onePort
		}
	}

	oneAddr := allAddrs[addr]
	oneAddr.PortMap = tmpMap
	allAddrs[addr] = oneAddr

	// log.Println("SAVEALL:", allAddrs)

	yaml.Write(appConfig.YamlPath, allAddrs)

	c.Redirect(http.StatusFound, "/scan/?addr="+addr)
}

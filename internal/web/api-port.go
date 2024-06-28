package web

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/aceberg/WatchYourPorts/internal/scan"
)

func apiPortScan(c *gin.Context) {

	addr := c.Param("addr")
	pStr := c.Param("port")

	pInt, err := strconv.Atoi(pStr)

	if err == nil {

		onePort := allAddrs[addr].PortMap[pInt]
		onePort.Port = pInt
		onePort.State = scan.IsOpen(addr, pInt)

		c.IndentedJSON(http.StatusOK, onePort)
	}
}

func apiAddrPortMap(c *gin.Context) {

	addr := c.Param("addr")

	c.IndentedJSON(http.StatusOK, allAddrs[addr].PortMap)
}

func apiHistory(c *gin.Context) {

	c.IndentedJSON(http.StatusOK, histAll)
}

func apiAllAddrs(c *gin.Context) {

	for key, value := range allAddrs {
		value.Total = 0
		value.Watching = 0
		value.Online = 0
		value.Offline = 0

		for _, p := range value.PortMap {
			value.Total++
			if p.Watch {
				value.Watching++
				if p.State {
					value.Online++
				} else {
					value.Offline++
				}
			}
		}

		allAddrs[key] = value
	}

	c.IndentedJSON(http.StatusOK, allAddrs)
}

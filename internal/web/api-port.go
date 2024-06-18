package web

import (
	// "log"
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

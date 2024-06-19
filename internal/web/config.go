package web

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/aceberg/WatchYourPorts/internal/check"
	"github.com/aceberg/WatchYourPorts/internal/conf"
	"github.com/aceberg/WatchYourPorts/internal/models"
)

func configHandler(c *gin.Context) {
	var guiData models.GuiData

	guiData.Config = appConfig

	guiData.Themes = []string{"cerulean", "cosmo", "cyborg", "darkly", "emerald", "flatly", "grass", "journal", "litera", "lumen", "lux", "materia", "minty", "morph", "pulse", "quartz", "sand", "sandstone", "simplex", "sketchy", "slate", "solar", "spacelab", "superhero", "united", "vapor", "yeti", "zephyr"}

	file, err := pubFS.ReadFile("public/version")
	check.IfError(err)
	version := string(file)
	guiData.Version = version[8:]

	c.HTML(http.StatusOK, "header.html", guiData)
	c.HTML(http.StatusOK, "config.html", guiData)
}

func saveConfigHandler(c *gin.Context) {

	appConfig.Host = c.PostForm("host")
	appConfig.Port = c.PostForm("port")
	appConfig.Theme = c.PostForm("theme")
	appConfig.Color = c.PostForm("color")
	tStr := c.PostForm("timeout")
	hStr := c.PostForm("trim")

	if tStr != "" {
		appConfig.Timeout, _ = strconv.Atoi(tStr)
	}
	if hStr != "" {
		appConfig.HistTrim, _ = strconv.Atoi(hStr)
	}

	conf.Write(appConfig)

	log.Println("INFO: writing new config to", appConfig.ConfPath)

	c.Redirect(http.StatusFound, "/config")
}

func saveInfluxHandler(c *gin.Context) {

	appConfig.InfluxAddr = c.PostForm("addr")
	appConfig.InfluxToken = c.PostForm("token")
	appConfig.InfluxOrg = c.PostForm("org")
	appConfig.InfluxBucket = c.PostForm("bucket")
	enable := c.PostForm("enable")
	skip := c.PostForm("skip")

	if enable == "on" {
		appConfig.InfluxEnable = true
	} else {
		appConfig.InfluxEnable = false
	}
	if skip == "on" {
		appConfig.InfluxSkipTLS = true
	} else {
		appConfig.InfluxSkipTLS = false
	}

	// conf.Write(appConfig)

	log.Println("INFO: writing new config to", appConfig)

	c.Redirect(http.StatusFound, "/config")
}

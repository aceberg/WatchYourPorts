package web

import (
	"html/template"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/aceberg/WatchYourPorts/internal/check"
	"github.com/aceberg/WatchYourPorts/internal/conf"
	"github.com/aceberg/WatchYourPorts/internal/models"
	"github.com/aceberg/WatchYourPorts/internal/yaml"
)

// Gui - start web server
func Gui(dirPath, nodePath string) {

	confPath := dirPath + "/config.yaml"
	check.Path(confPath)

	appConfig = conf.Get(confPath)

	appConfig.DirPath = dirPath
	appConfig.YamlPath = dirPath + "/hosts.yaml"
	appConfig.ConfPath = confPath
	appConfig.NodePath = nodePath

	// log.Println("INFO: starting web gui with config", appConfig.ConfPath)
	log.Println("INFO: starting web gui with config", appConfig)

	allAddrs = yaml.Read(appConfig.YamlPath)

	histAll = make(map[string]models.HistData)

	quitScan = make(chan bool)
	go routineScan(quitScan)

	address := appConfig.Host + ":" + appConfig.Port

	log.Println("=================================== ")
	log.Printf("Web GUI at http://%s", address)
	log.Println("=================================== ")

	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	templ := template.Must(template.New("").ParseFS(templFS, "templates/*"))
	router.SetHTMLTemplate(templ) // templates

	router.StaticFS("/fs/", http.FS(pubFS)) // public

	router.GET("/api/addr/:addr/", apiAddrPortMap)   // api-port.go
	router.GET("/api/port/:addr/:port", apiPortScan) // api-port.go

	router.GET("/", indexHandler)           // index.go
	router.GET("/config/", configHandler)   // config.go
	router.GET("/history/", historyHandler) // history.go
	router.GET("/scan/", scanHandler)       // scanpage.go

	router.POST("/addr_add/", addHandler)             // addr.go
	router.POST("/addr_del/", delHandler)             // addr.go
	router.POST("/addr_save/", renameHandler)         // addr.go
	router.POST("/config/", saveConfigHandler)        // config.go
	router.POST("/config_influx/", saveInfluxHandler) // config.go
	router.POST("/scan_save/", scanSaveHandler)       // scanpage.go

	err := router.Run(address)
	check.IfError(err)
}

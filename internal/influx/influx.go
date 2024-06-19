package influx

import (
	"context"
	"crypto/tls"
	"fmt"
	"log"
	"strings"

	"github.com/influxdata/influxdb-client-go/v2"

	"github.com/aceberg/WatchYourPorts/internal/check"
	"github.com/aceberg/WatchYourPorts/internal/models"
)

// Add - write data to InfluxDB2
func Add(appConfig models.Conf, oneHist models.HistData) {

	// client := influxdb2.NewClient(appConfig.InfluxAddr, appConfig.InfluxToken)

	client := influxdb2.NewClientWithOptions(appConfig.InfluxAddr, appConfig.InfluxToken,
		influxdb2.DefaultOptions().
			SetUseGZip(true).
			SetTLSConfig(&tls.Config{
				InsecureSkipVerify: true,
			}))

	writeAPI := client.WriteAPIBlocking(appConfig.InfluxOrg, appConfig.InfluxBucket)

	// Escape special characters in strings
	oneHist.Name = strings.ReplaceAll(oneHist.Name, " ", "\\ ")
	oneHist.Name = strings.ReplaceAll(oneHist.Name, ",", "\\,")
	oneHist.Name = strings.ReplaceAll(oneHist.Name, "=", "\\=")

	oneHist.PortName = strings.ReplaceAll(oneHist.PortName, " ", "\\ ")
	oneHist.PortName = strings.ReplaceAll(oneHist.PortName, ",", "\\,")
	oneHist.PortName = strings.ReplaceAll(oneHist.PortName, "=", "\\=")

	line := fmt.Sprintf("WatchYourPorts,addr=%s,port=%d,addrName=%s,portName=%s state=%t", oneHist.Addr, oneHist.Port, oneHist.Name, oneHist.PortName, oneHist.NowState)
	log.Println("LINE:", line)

	err := writeAPI.WriteRecord(context.Background(), line)
	check.IfError(err)

	client.Close()
}

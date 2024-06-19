package web

import (
	"log"
	"strconv"
	"time"

	"github.com/aceberg/WatchYourPorts/internal/influx"
	"github.com/aceberg/WatchYourPorts/internal/models"
	"github.com/aceberg/WatchYourPorts/internal/scan"
)

func routineScan(quit chan bool) {
	var lastDate time.Time

	if appConfig.Timeout == 0 {
		appConfig.Timeout = 10
	}

	for {
		select {
		case <-quit:
			return
		default:
			nowDate := time.Now()
			plusDate := lastDate.Add(time.Duration(appConfig.Timeout) * time.Minute)

			if nowDate.After(plusDate) {

				startScan()

				lastDate = time.Now()
			}

			time.Sleep(time.Duration(1) * time.Second)
		}
	}
}

func startScan() {
	var changed bool

	log.Println("Port scan started")

	for _, addr := range allAddrs {
		changed = false
		tmpPortMap := addr.PortMap

		for _, port := range addr.PortMap {

			if port.Watch {
				port.State = scan.IsOpen(addr.Addr, port.Port)
				tmpPortMap[port.Port] = port
				changed = true

				// History
				portStr := strconv.Itoa(port.Port)
				oneHist := histAll[addr.Addr+":"+portStr]
				oneHist.Name = addr.Name
				oneHist.Addr = addr.Addr
				oneHist.Port = port.Port
				oneHist.PortName = port.Name
				oneHist.NowState = port.State
				oneHist.State = append(oneHist.State,
					models.HistState{
						Date:  time.Now().Format("2006-01-02 15:04:05"),
						State: port.State,
					},
				)
				histAll[addr.Addr+":"+portStr] = oneHist

				if appConfig.InfluxEnable {
					influx.Add(appConfig, oneHist)
				}
			}
		}

		if changed {
			tmpAddr := addr
			tmpAddr.PortMap = tmpPortMap
			allAddrs[addr.Addr] = tmpAddr
		}
	}

	log.Println("Port scan done!")
}

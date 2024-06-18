package web

import (
	"log"
	"time"

	"github.com/aceberg/WatchYourPorts/internal/scan"
	// "github.com/aceberg/WatchYourPorts/internal/yaml"
)

func routineScan(quit chan bool) {
	var lastDate time.Time

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

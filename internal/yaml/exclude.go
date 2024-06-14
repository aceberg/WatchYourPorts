package yaml

import (
	"github.com/aceberg/WatchYourPorts/internal/models"
)

func ExcludeWatchFalse(ports map[int]models.PortItem) map[int]models.PortItem {

	cleanPorts := make(map[int]models.PortItem)

	for key, value := range ports {
		if value.Watch {
			cleanPorts[key] = value
		}
	}

	return cleanPorts
}

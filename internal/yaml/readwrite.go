package yaml

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"

	"github.com/aceberg/WatchYourPorts/internal/check"
	"github.com/aceberg/WatchYourPorts/internal/models"
)

// Read - read .yaml file to struct
func Read(path string) map[string]models.AddrToScan {

	file, err := os.ReadFile(path)
	check.IfError(err)

	plans := make(map[string]models.AddrToScan)

	err = yaml.Unmarshal(file, &plans)
	check.IfError(err)

	return plans
}

// Write - write struct to  .yaml file
func Write(path string, plans map[string]models.AddrToScan) {

	yamlData, err := yaml.Marshal(&plans)
	check.IfError(err)

	err = os.WriteFile(path, yamlData, 0644)
	check.IfError(err)

	log.Println("INFO: writing new plan file to", path)
}

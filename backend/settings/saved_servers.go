package settings

import (
	"github.com/Zaprit/ThePod/backend/server_repo"
	"github.com/labstack/gommon/log"
	"gopkg.in/yaml.v3"
	"os"
	"sync"
)

var loadedServers map[string]server_repo.Server
var runOnce sync.Once

func GetServers() map[string]server_repo.Server {
	runOnce.Do(loadServers)
	return loadedServers
}

func loadServers() {
	data, err := os.ReadFile(GetDataDir() + "/servers.yml")
	if err != nil {
		return
	}
	err = yaml.Unmarshal(data, &loadedServers)
	if err != nil {
		log.Warn("Failed to load server list from file")
	}
}

func SaveServers() {
	data, err := yaml.Marshal(&loadedServers)
	if err != nil {
		panic("Invalid servers??")
	}
	err = os.WriteFile(GetDataDir()+"/servers.yml", data, 0644)
	if err != nil {
		log.Warn("failed to write")
	}
}

//func AddServer() {
//
//}

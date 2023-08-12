package settings

import (
	"encoding/json"
	"github.com/labstack/gommon/log"
	"os"
	"sync"
)

//func (_ Servers) Get() []server_repo.Server {
//	runOnce.Do(loadServers)
//	return loadedServers
//}

var doOnce sync.Once

var loadedRepos []string

const defaultRepo = "https://zaprit.github.io/thepod"

func loadRepos() {
	data, err := os.ReadFile(GetDataDir() + "/repos.json")
	if err != nil { // Probably first run, IDK
		loadedRepos = []string{defaultRepo}
		go SaveRepos()
	}
	err = json.Unmarshal(data, &loadedRepos)
	if err != nil {
		log.Warn("Failed to load repo list from file")
	}
}

func GetRepos() {
	doOnce.Do(loadRepos)
}

func SaveRepos() {
	data, err := json.Marshal(&loadedServers)
	if err != nil {
		panic("Invalid repos???")
	}
	err = os.WriteFile(GetDataDir()+"/repos.json", data, 0644)
	if err != nil {
		log.Warn("failed to write")
	}
}

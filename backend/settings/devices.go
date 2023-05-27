package settings

import (
	"encoding/json"
	"github.com/Zaprit/ThePod/backend/patcher/remote"
	"github.com/google/uuid"
	"github.com/labstack/gommon/log"
	"os"
	"sync"
)

var doOnceDevices sync.Once

var loadedDevices map[uuid.UUID]remote.Device

func loadDevices() {
	data, err := os.ReadFile(GetDataDir() + "/repos.json")
	if err != nil {
		return
	}
	err = json.Unmarshal(data, &loadedRepos)
	if err != nil {
		log.Warn("Failed to load server list from file")
	}
}

func GetDevices() map[uuid.UUID]remote.Device {
	doOnceDevices.Do(loadDevices)
	return loadedDevices
}

func SaveDevices() {
	data, err := json.Marshal(&loadedDevices)
	if err != nil {
		panic("Invalid devices???")
	}
	err = os.WriteFile(GetDataDir()+"/devices.json", data, 0644)
	if err != nil {
		log.Warn("failed to write")
	}
}

func AddDevice(uuid uuid.UUID, device remote.Device) {
	loadedDevices[uuid] = device
}

package settings

import (
	"github.com/google/uuid"
	"github.com/labstack/gommon/log"
	"gopkg.in/yaml.v3"
	"os"
	"sync"
)

type Settings struct {
	RPCS3Path      string
	LastOpenServer int
	SelectedDevice uuid.UUID
}

var loadedSettings Settings

var loadSettingsOnce sync.Once

func loadSettings() {
	data, err := os.ReadFile(GetDataDir() + "/settings.yml")
	if err != nil {
		return
	}
	err = yaml.Unmarshal(data, &loadedSettings)
	if err != nil {
		log.Warn("Failed to load settings from file")
	}
}

func SaveSettings() {
	data, err := yaml.Marshal(&loadedSettings)
	if err != nil {
		panic("Invalid servers??")
	}
	err = os.WriteFile(GetDataDir()+"/settings.yml", data, 0644)
	if err != nil {
		log.Warn("failed to write")
	}
}

func GetSettings() *Settings {
	loadSettingsOnce.Do(loadSettings)
	return &loadedSettings
}

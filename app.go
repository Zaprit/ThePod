package main

import (
	"context"
	"github.com/Zaprit/ThePod/backend/patcher"
	"github.com/Zaprit/ThePod/backend/patcher/remote"
	"github.com/Zaprit/ThePod/backend/settings"
	"github.com/google/uuid"
	"github.com/wailsapp/wails/v2/pkg/runtime"
	"os"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so that we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
	runtime.WindowSetDarkTheme(ctx)
}

func (a *App) AddDevice(name string, hostname string, port uint16, deviceType remote.DeviceType) {
	device := remote.Device{
		Name:     name,
		Hostname: hostname,
		Port:     port,
		Type:     deviceType,
	}
	settings.AddDevice(uuid.New(), device)
	settings.SaveDevices()
}

func (a *App) PatchDevice(uuidString string, url string) error {
	id, err := uuid.Parse(uuidString)
	if err != nil {
		return err
	}
	device := settings.GetDevices()[id]
	return remote.Patch(device, url)
}

func (a *App) PatchFile(path string, url string, outFilePath string) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	foundURLs, err := patcher.FindURLs(data)
	if err != nil {
		return err
	}
	for _, index := range foundURLs {
		patcher.Patch(data, index, url)
	}
	return os.WriteFile(outFilePath, data, 0644)
}

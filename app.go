package main

import (
	"context"
	"fmt"
	"github.com/Zaprit/ThePod/backend/img_cache"
	"github.com/Zaprit/ThePod/backend/news"
	"github.com/Zaprit/ThePod/backend/patcher"
	"github.com/Zaprit/ThePod/backend/patcher/remote"
	"github.com/Zaprit/ThePod/backend/patcher/remote/auto_detect"
	"github.com/Zaprit/ThePod/backend/server_repo"
	"github.com/Zaprit/ThePod/backend/settings"
	"github.com/google/uuid"
	"github.com/labstack/gommon/log"
	"github.com/wailsapp/wails/v2/pkg/runtime"
	"os"
)

// App struct
type App struct {
	ctx            context.Context
	selectedServer *server_repo.Server
	selectedDevice *remote.Device
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
	go auto_detect.AutoDetect() // Start auto-detecting for PS3s
	// This is awful, find better way to do it
	var servers []server_repo.Server
	for _, server := range a.GetServers() {
		servers = append(servers, server)
	}
	a.selectedServer = &servers[0]
}

func (a *App) shutdown(_ context.Context) {
	settings.SaveServers()
	settings.SaveRepos()
	settings.SaveDevices()
	settings.SaveSettings()
}

func (a *App) AddDevice(name string, hostname string, port uint16, deviceType remote.Platform) {
	device := &remote.Device{
		Name:     name,
		Hostname: hostname,
		Port:     port,
		Type:     deviceType,
	}
	settings.AddDevice(uuid.New(), device)
	settings.SaveDevices()
}

func (a *App) PatchDevice(uuidString string, url, titleID string) error {
	id, err := uuid.Parse(uuidString)
	if err != nil {
		return err
	}
	device := settings.GetDevices()[id]
	return remote.Patch(device, url, titleID)
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

func (a *App) GetServer(id string) server_repo.Server {
	return settings.GetServers()[id]
}

func (a *App) GetServers() map[string]server_repo.Server {
	return settings.GetServers()
}

func (a *App) SelectServer(id string) {
	server := settings.GetServers()[id]
	a.selectedServer = &server
	runtime.LogDebugf(a.ctx, "Changed server to %s", id)
	runtime.EventsEmit(a.ctx, "server_change", server)
}

func (a *App) GetSelectedServer() *server_repo.Server {
	return a.selectedServer
}

func (a *App) AddServer(server server_repo.Server) {
	settings.AddServer(server)
}

func (a *App) GetArticles() []news.Item {
	feed := news.GetArticles(a.ctx, a.selectedServer.FeedURL)
	return feed
}

func (a *App) GetDevices() map[uuid.UUID]*remote.Device {
	return settings.GetDevices()
}

func (a *App) GetDevice(uuidStr string) *remote.Device {
	fmt.Println(uuidStr)
	if uuidStr == "" {
		return nil
	}
	return settings.GetDevices()[uuid.MustParse(uuidStr)]
}

func (a *App) GetSelectedDevice() *remote.Device {
	return a.selectedDevice
}

func (a *App) FetchImage(url string) string {
	cacheURL, err := img_cache.DownloadImage(url)
	if err != nil {
		log.Error(err)
	}
	return cacheURL
}

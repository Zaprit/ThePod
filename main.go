package main

import (
	"embed"
	"github.com/Zaprit/ThePod/backend/news"
	"github.com/Zaprit/ThePod/backend/settings"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"github.com/wailsapp/wails/v2/pkg/options/mac"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	// Create an instance of the app structure
	app := NewApp()

	appOptions := &options.App{
		Title: "ThePod",
		Mac: &mac.Options{
			About: &mac.AboutInfo{
				Title:   "ThePod",
				Message: "A LittleBigPlanet companion app for community custom servers",
				Icon:    nil,
			},
			Appearance: mac.NSAppearanceNameDarkAqua,
			//TitleBar: &mac.TitleBar{TitlebarAppearsTransparent: true, HideTitle: true}, // Assess whether or not this will work
		},
		Width:  1024,
		Height: 768,
		//Frameless: true,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup:        app.startup,
		Bind: []interface{}{
			app,
			&news.ArticleFetcher{},
		},
	}

	//switch runtime.GOOS {
	//case "darwin":
	//	appOptions.Frameless = false
	//}

	// Create application with options
	err := wails.Run(appOptions)
	if err != nil {
		println("Error:", err.Error())
	}

	settings.SaveServers()
	settings.SaveRepos()
	settings.SaveDevices()
	settings.SaveSettings()
}

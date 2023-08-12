package main

import (
	"embed"
	"github.com/Zaprit/ThePod/backend/img_cache"
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
			Assets:  assets,
			Handler: img_cache.NewImgCache(),
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup:        app.startup,
		OnShutdown:       app.shutdown,
		Bind: []interface{}{
			app,
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
}

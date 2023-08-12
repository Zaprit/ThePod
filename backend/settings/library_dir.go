package settings

import (
	"os"
	"runtime"
)

func getAppdataDir() string {
	switch runtime.GOOS {
	case "windows":
		return os.Getenv("APPDATA") + "/ThePod"
	case "darwin":
		return os.Getenv("HOME") + "/Library/Application Support/ThePod"
	default:
		return os.Getenv("HOME") + "/.local/share/ThePod"
	}
}

func getCacheDir() string {
	switch runtime.GOOS {
	case "windows":
		return os.Getenv("TEMP") + "/ThePod"
	case "darwin":
		return os.Getenv("HOME") + "/Library/Caches/ThePod"
	default:
		return os.Getenv("HOME") + "/.cache/ThePod"
	}
}

func GetDataDir() string {
	dir := getAppdataDir()
	err := os.MkdirAll(dir, 0750)
	if err != nil {
		panic("Failed to create data dirs")
	}
	return dir
}

func GetCacheDir(subdir string) string {
	dir := getCacheDir() + "/" + subdir
	err := os.MkdirAll(dir, 0750)
	if err != nil {
		panic("Failed to create data dirs")
	}
	return dir
}

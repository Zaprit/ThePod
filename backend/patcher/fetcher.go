package patcher

import (
	"errors"
	"github.com/Zaprit/ThePod/backend/settings"
	"io"
	"net/http"
	"os"
	"runtime"
)

// Gets various tools we might need for patching

// PreparePatchingEnvironment creates all the folders we will need for patching, downloads all the tools/keys and
func PreparePatchingEnvironment() error {
	cacheDir := settings.GetCacheDir("patcher")
	err := os.Mkdir(cacheDir+"/data", 0755)
	if err != nil {
		return err
	}

	err = getSCETool(cacheDir + "/scetool")

	err = GetKeys(cacheDir + "/data/keys")
	if err != nil {
		return err
	}
	err = GetLDRCurves(cacheDir + "/data/ldr_curves")
	if err != nil {
		return err
	}
	err = GetVSHCurves(cacheDir + "/data/vsh_curves")
	if err != nil {
		return err
	}

	return nil
}

const (
	SCEToolDownloadURL = "https://github.com/Zaprit/ThePod_extras/releases/download/latest/" // TODO: Fill this in, after creating a supporting files repo
)

func getSCEToolFileName() (string, error) {
	switch runtime.GOOS {
	case "windows":
		return "scetool_" + runtime.GOARCH + ".exe", nil
	case "darwin":
		return "scetool_darwin_universal", nil
	case "linux":
		return "scetool_linux_" + runtime.GOARCH, nil
	default:
		return "", errors.New("unsupported platform")
	}
}

func getSCETool(outputPath string) error {
	filename, err := getSCEToolFileName()
	if err != nil {
		return err
	}
	url := SCEToolDownloadURL + filename

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}
	defer req.Body.Close()

	file, err := os.OpenFile(outputPath, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		return err
	}

	_, err = io.Copy(file, req.Body)

	return err
}

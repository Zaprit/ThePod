package patcher

import (
	"github.com/labstack/gommon/log"
	"io"
	"net/http"
	"os"
)

// We can't distribute keys, so we'll go and get them from elsewhere
// Whatever it was sony, I didn't do it

const (
	keysURL      = "https://raw.githubusercontent.com/rhynec/PS3-Keys/main/keys"
	vshCurvesURL = "https://raw.githubusercontent.com/rhynec/PS3-Keys/main/vsh_curves"
	ldrCurvesURL = "https://raw.githubusercontent.com/rhynec/PS3-Keys/main/ldr_curves"
)

func GetKeys(outputPath string) error {
	req, err := http.NewRequest("GET", keysURL, nil)
	if err != nil {
		return err
	}
	file, err := os.OpenFile(outputPath, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		return err
	}
	_, err = io.Copy(file, req.Body)
	return err
}
func GetVSHCurves(outputPath string) error {
	req, err := http.NewRequest("GET", vshCurvesURL, nil)
	if err != nil {
		return err
	}
	file, err := os.OpenFile(outputPath, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		return err
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Error("Failed to close http body")
		}
	}(req.Body)

	_, err = io.Copy(file, req.Body)
	return err
}
func GetLDRCurves(outputPath string) error {
	req, err := http.NewRequest("GET", ldrCurvesURL, nil)
	if err != nil {
		return err
	}
	file, err := os.OpenFile(outputPath, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		return err
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Error("Failed to close http body")
		}
	}(req.Body)
	_, err = io.Copy(file, req.Body)
	return err
}

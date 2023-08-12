package img_cache

import (
	"github.com/Zaprit/ThePod/backend/settings"
	"github.com/google/uuid"
	"io"
	"net/http"
	"os"
)

func DownloadImage(url string) (string, error) {
	cacheDir := settings.GetCacheDir("imgCache")
	err := os.MkdirAll(cacheDir, 0750)
	if err != nil {
		return "", err
	}

	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	imgName := uuid.New().String()

	file, err := os.OpenFile(cacheDir+"/"+imgName, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return "", err
	}

	_, err = io.Copy(file, resp.Body)
	if err != nil {
		return "", err
	}

	err = resp.Body.Close()
	return imgName, err
}

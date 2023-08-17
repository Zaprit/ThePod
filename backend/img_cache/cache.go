package img_cache

import (
	"github.com/Zaprit/ThePod/backend/settings"
	"hash/crc32"
	"io"
	"net/http"
	"os"
	"strconv"
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

	imgBytes, err := io.ReadAll(resp.Body)
	imgName := strconv.FormatInt(int64(crc32.ChecksumIEEE(imgBytes)), 16)
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

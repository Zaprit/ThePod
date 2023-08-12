package img_cache

import (
	"fmt"
	"github.com/Zaprit/ThePod/backend/settings"
	"github.com/labstack/gommon/log"
	"io"
	"net/http"
	"os"
	"strings"
)

type ImgCache struct {
	http.Handler
}

func NewImgCache() *ImgCache {
	return &ImgCache{}
}

func (h *ImgCache) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	cacheDir := settings.GetCacheDir("imgCache")

	var err error
	requestedFilename := strings.TrimPrefix(req.URL.Path, "/")

	println("Requesting file:", cacheDir+requestedFilename)
	file, err := os.Open(cacheDir + requestedFilename)
	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		_, _ = res.Write([]byte(fmt.Sprintf("Could not load file %s", requestedFilename)))
	}

	_, err = io.Copy(res, file)
	if err != nil {
		log.Error(err.Error())
	}
}

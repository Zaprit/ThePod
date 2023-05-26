package server_repo

import (
	"github.com/labstack/gommon/log"
	"net/http"
)

func GetServers(repoURL string) []Server {
	req, err := http.NewRequest("GET", repoURL, nil)
	if err != nil {
		log.Error("failed to get repo, ", err.Error())
		return nil
	}

	http.DefaultClient.Do(req)
	return nil
}

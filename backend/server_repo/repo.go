package server_repo

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const repoVersion = "1.0"

func GetServerList(repoURL string) ([]string, error) {
	req, err := http.NewRequest("GET", repoURL+"/podrepo.json", nil)
	if err != nil {
		return nil, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to get server, response from server was %d %s", resp.StatusCode, http.StatusText(resp.StatusCode))
	}

	defer resp.Body.Close()

	repoBytes, err := io.ReadAll(resp.Body)
	var repo Repo

	err = json.Unmarshal(repoBytes, &repo)
	if err != nil {
		return nil, err
	}

	if repo.Version != repoVersion {
		return nil, OutDatedClientError("repo version from server was \"%s\", however we can only read \"%s\"", repo.Version, repoVersion)
	}

	return repo.Servers, nil
}

func GetServer(repoURL, serverName string) (*Server, error) {
	req, err := http.NewRequest("GET", repoURL+"/servers/"+serverName+".json", nil)
	if err != nil {
		return nil, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	serverBytes, err := io.ReadAll(resp.Body)
	server := new(Server)
	err = json.Unmarshal(serverBytes, server)
	if err != nil {
		return nil, err
	}
	return server, nil
}

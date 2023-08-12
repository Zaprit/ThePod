package lighthouse

import (
	"encoding/json"
	"github.com/Zaprit/ThePod/backend/api_client/lighthouse/model"
	"github.com/Zaprit/ThePod/backend/server_repo"
	"github.com/labstack/gommon/log"
	"io"
	"net/http"
)

func V1LighthouseApiGetUser(server server_repo.Server, userID string) (*model.User, error) {
	resp, err := http.Get(server.APIUrl + "/user/" + userID)
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	err = resp.Body.Close()
	if err != nil {
		log.Warn("Failed to close body", err)
	}

	user := new(model.User)
	err = json.Unmarshal(body, &user)

	return user, err
}

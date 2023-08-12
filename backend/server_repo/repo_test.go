package server_repo

import (
	"fmt"
	"testing"
)

func TestGetServers(t *testing.T) {
	servers, err := GetServerList("https://zaprit.github.io/thepod")
	if err != nil {
		t.Error(err)
	}
	for _, server := range servers {
		server, err := GetServer("https://zaprit.github.io/thepod", server)
		if err != nil {
			t.Error(err)
		}
		fmt.Println(server.DisplayName)
		fmt.Println(server.Description)
		fmt.Println(server.Owner)
	}

}

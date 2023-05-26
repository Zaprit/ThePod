package rpcs3_patcher

import (
	"fmt"
	"testing"
)

func TestCreatePatch(t *testing.T) {
	fmt.Println(CreatePatch(URLBundle{
		HttpsURL:    "https://lbpunion.com",
		HttpURL:     "https://google.com",
		PresenceURL: "https://valtek.uk",
		LiveURL:     "https://youtube.com",
	}))
}

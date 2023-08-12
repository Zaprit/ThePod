package remote

import (
	"errors"
	"fmt"
	"github.com/jlaffaye/ftp"
	"github.com/labstack/gommon/log"
	"strings"
	"time"
)

func Patch(device *Device, url string, gameID string) error {
	switch device.Type {
	case PS3:
		return remotePatchPS3(device, url, gameID)
	default:
		return errors.New("unsupported device type")
	}
}

func remotePatchPS3(device *Device, url string, gameID string) error {
	c, err := ftp.Dial(fmt.Sprintf("%s:%d", device.Hostname, device.Port), ftp.DialWithTimeout(5*time.Second))
	if err != nil {
		return err
	}
	err = c.Login("anonymous", "anonymous")
	if err != nil {
		return err
	}

	ents, err := c.List("/dev_hdd0/game")
	if err != nil {
		log.Error(err.Error())
	}
	for _, ent := range ents {
		if !strings.HasPrefix(ent.Name, "B") && !strings.HasPrefix(ent.Name, "N") {
			continue
		}
		game, exists := GameIDs[ent.Name]
		if !exists {
			continue
		}

		fmt.Println(game)

		fmt.Println(ent.Name)
	}

	return nil
}

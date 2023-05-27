package remote

import "errors"

func Patch(device Device, url string) error {
	switch device.Type {
	case PS3:
		return remotePatchPS3(device, url)
	default:
		return errors.New("unsupported device type")
	}
}

func remotePatchPS3(device Device, url string) error {
	return nil
}

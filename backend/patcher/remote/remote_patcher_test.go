package remote

import "testing"

func TestPatch(t *testing.T) {
	dev := Device{
		Name:     "Test",
		Hostname: "192.168.178.233",
		Port:     21,
		Type:     PS3,
	}
	err := Patch(&dev, "http://dev.lbp.valtek.local/api/LBP_XML", "BCES00850")
	if err != nil {
		t.Error(err)
	}
}

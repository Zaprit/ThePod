package remote

type DeviceType int

type Device struct {
	Name     string   `json:"name,omitempty"`
	Hostname string   `json:"hostname,omitempty"`
	Port     uint16   `json:"port,omitempty"`
	Type     Platform `json:"type,omitempty"`
}

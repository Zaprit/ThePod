package remote

type DeviceType int

const (
	PS3 DeviceType = iota
	PSVita
	//PSP // Unused for a reason, ain't nobody wanting to patch their psp for lbp... yet
)

type Device struct {
	Name     string
	Hostname string
	Port     uint16
	Type     DeviceType
}

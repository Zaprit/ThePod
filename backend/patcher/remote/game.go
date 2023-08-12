package remote

type Game uint8
type Platform uint8

const (
	LittleBigPlanet Game = iota
	LittleBigPlanetBeta
	LittleBigPlanetPSP // Currently Unused
	LittleBigPlanet2
	LittleBigPlanet2Beta
	LittleBigPlanetVita
	LittleBigPlanet3
	LittleBigPlanet3Beta
)

const (
	PS3 Platform = iota
	PSP
	PSVita
	PS4
)

type EBoot struct {
	TitleID  string
	Game     Game
	Platform Platform
}

package patcher

type Game uint8
type Platform uint8

const (
	LittleBigPlanet    Game = iota
	LittleBigPlanetPSP      // Currently Unused
	LittleBigPlanet2
	LittleBigPlanetVita
	LittleBigPlanet3
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

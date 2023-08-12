package server_repo

import "fmt"

// Server represents one lbp custom server, the name should be globally unique, per server. This will likely not happen but one can hope
// Icons will be saved by the icon handler and butchered into something more cached
type Server struct {
	Name            string         `json:"name"`
	Website         string         `json:"website"`
	GameUrl         string         `json:"game_url"`
	APIUrl          string         `json:"api_url,omitempty"`
	FeedURL         string         `json:"feed_url,omitempty"`
	OpenAuth        bool           `json:"open_auth"`
	IconUrl         string         `json:"icon_url,omitempty"`
	ModsAllowed     bool           `json:"mods_allowed"`
	DisplayName     string         `json:"display_name"`
	Description     string         `json:"description"`
	Owner           string         `json:"owner"`
	AccentColour    string         `json:"accent_colour"`
	BackgroundImage string         `json:"background_image,omitempty"`
	SupportedGames  SupportedGames `json:"supported_games,omitempty"`
}

type SupportedGames struct {
	Lbp1   bool `json:"lbp1"`
	Lbp2   bool `json:"lbp2"`
	Lbp3   bool `json:"lbp3"`
	LbpV   bool `json:"lbp_vita"`
	LbpPSP bool `json:"lbp_psp"`
}

type Repo struct {
	Version string   `json:"version"`
	Servers []string `json:"servers"`
}

func OutDatedClientError(format string, args ...string) OutDatedClient {
	return OutDatedClient{message: fmt.Sprintf(format, args)}
}

type OutDatedClient struct {
	message string
}

func (c OutDatedClient) Error() string {
	return c.message
}

package server_repo

// Server represents one lbp custom server, the name should be globally unique, per server. This will likely not happen but one can hope
type Server struct {
	Name           string `json:"name"`
	Website        string `json:"website"`
	GameUrl        string `json:"game_url"`
	OpenAuth       bool   `json:"open_auth"`
	IconUrl        string `json:"icon_url"`
	ModsAllowed    bool   `json:"mods_allowed"`
	DisplayName    string `json:"display_name"`
	Description    string `json:"description"`
	Owner          string `json:"owner"`
	AccentColour   string `json:"accent_colour"`
	SupportedGames struct {
		Lbp1   bool `json:"lbp1"`
		Lbp2   bool `json:"lbp2"`
		Lbp3   bool `json:"lbp3"`
		LbpV   bool `json:"lbpv"`
		LbpPSP bool `json:"lbppsp"`
	} `json:"supported_games"`
}

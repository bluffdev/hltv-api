package handlers

type Country struct {
	Name string `json:"name"`
	Flag string `json:"flag"`
}

// Used with team and teams
type Player struct {
	Fullname string  `json:"fullname"`
	Image    string  `json:"image"`
	Nickname string  `json:"nickname"`
	Country  Country `json:"country"`
}

type Team struct {
	Id      int      `json:"id"`
	Ranking int      `json:"ranking"`
	Name    string   `json:"name"`
	Logo    string   `json:"logo"`
	Players []Player `json:"players"`
}

// Used with player and players
type PlayerStats struct {
	Id         int    `json:"id"`
	Nickname   string `json:"nickname"`
	Team       string `json:"team"`
	Slug       string `json:"slug"`
	MapsPlayed string `json:"mapsPlayed"`
	Kd         string `json:"kd"`
	Rating     string `json:"rating"`
}
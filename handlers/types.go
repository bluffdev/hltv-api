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

// Used with players.go
type PlayerStats struct {
	Id         int    `json:"id"`
	Nickname   string `json:"nickname"`
	Team       string `json:"team"`
	Slug       string `json:"slug"`
	MapsPlayed string `json:"mapsPlayed"`
	Kd         string `json:"kd"`
	Rating     string `json:"rating"`
}

// Used with player.go
type PlayerTeam struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

// Used with player.go
type PlayerStats2 struct {
	Id           int        `json:"id"`
	Team         PlayerTeam `json:"team"`
	Image        string     `json:"image"`
	Nickname     string     `json:"nickname"`
	Age          int        `json:"age"`
	Rating       float64    `json:"rating"`
	Impact       float64    `json:"impact"`
	Dpr          float64    `json:"dpr"`
	Apr          float64    `json:"apr"`
	Kast         float64    `json:"kast"`
	Kpr          float64    `json:"kpr"`
	HsPercentage float64    `json:"hsPercentage"`
	MapsPlayed   int        `json:"mapsPlayed"`
}

// Used with matches
type Event struct {
	Name string `json:"name"`
	Logo string `json:"logo"`
}

type MatchTeam struct {
	Name string `json:"name"`
	Logo string `json:"logo"`
}

type Match struct {
	Id    int         `json:"id"`
	Date  string      `json:"date"`
	Time  string      `json:"time"`
	Event Event       `json:"event"`
	Stars int         `json:"stars"`
	Maps  string      `json:"maps"`
	Teams []MatchTeam `json:"teams"`
}

type ResultTeam struct {
	Name  string `json:"name"`
	Logo  string `json:"logo"`
	Score int    `json:"score"`
}

type Result struct {
	Event   Event        `json:"event"`
	Maps    string       `json:"maps"`
	Date    string       `json:"date"`
	Teams   []ResultTeam `json:"teams"`
	MatchId int          `json:"matchId"`
}

type Stats struct {
	Name     string  `json:"name"`
	Nickname string  `json:"nickname"`
	Id       int     `json:"id"`
	Kills    int     `json:"kills"`
	Deaths   int     `json:"deaths"`
	Adr      float64 `json:"adr"`
	Kast     float64 `json:"kast"`
	Rating   float64 `json:"rating"`
}

type TeamStats struct {
	Name    string  `json:"name"`
	Logo    string  `json:"logo"`
	Result  int     `json:"result"`
	Players []Stats `json:"players"`
}

type Match2 struct {
	Id    int         `json:"id"`
	Date  string      `json:"date"`
	Teams []TeamStats `json:"teams"`
}

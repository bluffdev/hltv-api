package handlers

type Country struct {
	Name string
	Flag string
}

type Player struct {
	Fullname string
	Image    string
	Nickname string
	Country  Country
}

type Team struct {
	Id      int
	Ranking int
	Name    string
	Logo    string
	Players []Player
}
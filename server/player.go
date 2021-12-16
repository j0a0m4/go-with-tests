package server

type Player struct {
	Name string `json:"name,omitempty"`
	Wins int    `json:"wins,omitempty"`
}

// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Player struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Position string `json:"position"`
}

type Team struct {
	ID           string    `json:"id"`
	Name         string    `json:"name"`
	Location     *string   `json:"location"`
	Players      []*Player `json:"players"`
	Conference   string    `json:"conference"`
	Abbreviation string    `json:"abbreviation"`
}

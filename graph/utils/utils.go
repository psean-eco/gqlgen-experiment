package graph

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/psean/gql-go-gen/graph/model"
)

type PlayerResponse struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Position  string `json:"position"`
}

type TeamsResponse struct {
	ID           int    `json:"id"`
	Abbreviation string `json:"abbreviation"`
	City         string `json:"city"`
	Conference   string `json:"conference"`
	Division     string `json:"division"`
	FullName     string `json:"full_name"`
	Name         string `json:"name"`
}

type MetaResponse struct {
	TotalPages  int `json:"total_pages"`
	CurrentPage int `json:"current_page"`
	NextPage    int `json:"next_page"`
	PerPage     int `json:"per_page"`
	TotalCount  int `json:"total_count"`
}

type Response struct {
	Data []PlayerResponse `json:"data"`
}

type TeamResponse struct {
	Data []TeamsResponse `json:"data"`
}

func GetPlayers() ([]*model.Player, error) {
	response, err := http.Get("https://www.balldontlie.io/api/v1/players")
	player := model.Player{
		ID:       "1234",
		Name:     "test Device1",
		Position: "C",
	}
	var defaultPlayers []*model.Player
	defaultPlayers = append(defaultPlayers, &player)
	if err != nil {
		log.Fatal(err)
		return defaultPlayers, err
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
		return defaultPlayers, err
	}
	if err != nil {
		log.Fatal(err)
		return defaultPlayers, err
	}
	responseObject := new(Response)
	json.Unmarshal(responseData, &responseObject)

	var players []*model.Player

	for _, player := range responseObject.Data {
		var NewPlayer = model.Player{
			ID:       strconv.Itoa(player.ID),
			Name:     player.LastName,
			Position: player.Position,
		}
		players = append(players, &NewPlayer)
	}

	return players, err

}
func GetTeams() ([]*model.Team, error) {
	response, err := http.Get("https://www.balldontlie.io/api/v1/teams")
	var lo = "home"
	team := model.Team{
		ID:       "1234",
		Name:     "test Device1",
		Location: &lo,
	}
	var defaultTeams []*model.Team
	defaultTeams = append(defaultTeams, &team)
	if err != nil {
		log.Fatal(err)
		return defaultTeams, err
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
		return defaultTeams, err
	}
	if err != nil {
		log.Fatal(err)
		return defaultTeams, err
	}
	responseObject := new(TeamResponse)
	json.Unmarshal(responseData, &responseObject)

	var teams []*model.Team

	for _, team := range responseObject.Data {
		var NewTeam = model.Team{
			ID:   strconv.Itoa(team.ID),
			Name: team.Name,
		}
		teams = append(teams, &NewTeam)
	}

	return teams, err

}

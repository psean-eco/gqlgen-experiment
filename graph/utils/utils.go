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
	ID        int           `json:"id"`
	FirstName string        `json:"first_name"`
	LastName  string        `json:"last_name"`
	Position  string        `json:"position"`
	Team      TeamsResponse `json:"team"`
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
	var players []*model.Player
	if err != nil {
		log.Fatal(err)
		return players, err
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
		return players, err
	}
	if err != nil {
		log.Fatal(err)
		return players, err
	}
	responseObject := new(Response)
	json.Unmarshal(responseData, &responseObject)

	for _, player := range responseObject.Data {
		newPlayer := convertToPlayer(player)
		players = append(players, &newPlayer)
	}

	return players, err

}
func GetTeams() ([]*model.Team, error) {
	response, err := http.Get("https://www.balldontlie.io/api/v1/teams")
	var teams []*model.Team
	if err != nil {
		log.Fatal(err)
		return teams, err
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
		return teams, err
	}
	if err != nil {
		log.Fatal(err)
		return teams, err
	}
	responseObject := new(TeamResponse)
	json.Unmarshal(responseData, &responseObject)

	for _, team := range responseObject.Data {
		convertedTeam := convertToTeam(team)
		teams = append(teams, &convertedTeam)
	}

	return teams, err

}

func GetPlayersForTeam(team model.Team) ([]*model.Player, error) {
	response, err := http.Get("https://www.balldontlie.io/api/v1/players")
	var players []*model.Player
	if err != nil {
		log.Fatal(err)
		return players, err
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
		return players, err
	}
	if err != nil {
		log.Fatal(err)
		return players, err
	}
	responseObject := new(Response)
	json.Unmarshal(responseData, &responseObject)

	for _, player := range responseObject.Data {
		if strconv.Itoa(player.Team.ID) == team.ID {
			newPlayer := convertToPlayer(player)
			players = append(players, &newPlayer)
		}
	}

	return players, err

}
func convertToTeam(team TeamsResponse) model.Team {
	newTeam := model.Team{
		ID:       strconv.Itoa(team.ID),
		Name:     team.Name,
		Location: &team.City,
	}
	return newTeam
}

func convertToPlayer(player PlayerResponse) model.Player {
	newPlayer := model.Player{
		ID:       strconv.Itoa(player.ID),
		Name:     player.FirstName + " " + player.LastName,
		Position: player.Position,
	}
	return newPlayer
}

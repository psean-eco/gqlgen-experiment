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
	id        string `json:"id"`
	position  string `json:"position"`
	firstName string `json:"first_name"`
}

type Response struct {
	Data []struct {
		ID           int         `json:"id"`
		FirstName    string      `json:"first_name"`
		HeightFeet   interface{} `json:"height_feet"`
		HeightInches interface{} `json:"height_inches"`
		LastName     string      `json:"last_name"`
		Position     string      `json:"position"`
		Team         struct {
			ID           int    `json:"id"`
			Abbreviation string `json:"abbreviation"`
			City         string `json:"city"`
			Conference   string `json:"conference"`
			Division     string `json:"division"`
			FullName     string `json:"full_name"`
			Name         string `json:"name"`
		} `json:"team"`
		WeightPounds interface{} `json:"weight_pounds"`
	} `json:"data"`
	Meta struct {
		TotalPages  int `json:"total_pages"`
		CurrentPage int `json:"current_page"`
		NextPage    int `json:"next_page"`
		PerPage     int `json:"per_page"`
		TotalCount  int `json:"total_count"`
	} `json:"meta"`
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



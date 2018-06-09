package hackathon_api

import "errors"

type Challenge struct {
	Id     string `json:"id" bson:"_id"`
	Name   string `json:"name"`
	Points int    `json:"points"`
}

var Challenges []Challenge

func init() {
	Challenges = []Challenge{
		{
			"install_solarpanels",
			"Zonnepanelen installeren",
			100,
		},
	}
}

func GetChallenge(id string) (*Challenge, error) {
	for _, challenge := range Challenges {
		if challenge.Id == id {
			return &challenge, nil
		}
	}

	return nil, errors.New("not found")
}

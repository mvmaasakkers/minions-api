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
		{"zonnecollectoren", "Zonnecollectoren", 350},
		{"energielabel", "Energielabel", 100},
		{"spouwmuur", "Spouwmuur", 500},
		{"vloer_bodem_isolatie", "Vloer bodem isolatie", 350},
		{"dakisolatie", "Dakisolatie", 500},
		{"hr_ketel", "HR ketel", 250},
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

package hackathon_api

import "errors"

type Challenge struct {
	Id     string `json:"id" bson:"_id"`
	Name   string `json:"name"`
	Points int    `json:"points"`
	Done   bool   `json:"done"`
}

var Challenges []Challenge

func init() {
	Challenges = []Challenge{
		{"zonnecollectoren", "Zonnecollectoren", 350, false},
		{"energielabel", "Energielabel", 100, false},
		{"spouwmuur", "Spouwmuur", 500, false},
		{"vloer_bodem_isolatie", "Vloer bodem isolatie", 350, false},
		{"dakisolatie", "Dakisolatie", 500, false},
		{"hr_ketel", "HR ketel", 250, false},
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

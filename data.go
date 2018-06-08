package hackathon_api

import (
	"time"
	"gopkg.in/mgo.v2/bson"
	"encoding/json"
)

type Data struct {
	Id   bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Time time.Time     `json:"time"`

	Source DataSource `json:"source"`
	Type   string     `json:"type" validate:"nonzero"`
	Event  string     `json:"event" validate:"nonzero"`

	Data json.RawMessage `json:"data"`
}

type DataSource struct {
	ID string `json:"id" validate:"nonzero"`
	Type string `json:"type" validate:"nonzero"`
}
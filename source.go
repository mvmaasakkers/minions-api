package hackathon_api

import "gopkg.in/mgo.v2/bson"

type Source struct {
	Id bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Type string `json:"type" validate:"nonzero"`
	Name string `json:"name" validate:"nonzero"`
	Meta map[string]interface{} `json:"meta"`
}

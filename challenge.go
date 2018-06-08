package hackathon_api

import "gopkg.in/mgo.v2/bson"

type Challenge struct {
	Id bson.ObjectId `json:"id" bson:"_id"`
	Name string `json:"name"`
}

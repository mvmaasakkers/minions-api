package hackathon_api

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

type Token struct {
	Id        bson.ObjectId `json:"id" bson:"_id,omitempty"`
	CreatedAt time.Time     `json:"created_at"`
	ExpiresAt time.Time     `json:"expires_at"`
	Token     string        `json:"token"`
	UserId    string        `json:"user_id"`
}

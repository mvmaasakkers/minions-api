package hackathon_api

import "gopkg.in/mgo.v2/bson"

type User struct {
	Id       bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Username string        `json:"username"`
	Password string        `json:"password,omitempty"`
	Email    string        `json:"email"`

	BankUsers []BankUser `json:"bankusers"`
}

type BankUser struct {
	Username string `json:"username" validate:"nonzero"`
	Password string `json:"password" validate:"nonzero"`
	Email    string `json:"email" validate:"nonzero"`
}

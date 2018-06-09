package mongodb

import (
	"gopkg.in/mgo.v2"
)

type ChallengeService struct {
	DB *DB
	Collection string
	database *mgo.Database
}

func NewChallengeService(db *DB) *ChallengeService {
	svc := &ChallengeService{DB: db, Collection: "challenge", database: db.Session.DB(db.Settings.Database)}

	return svc
}
package mongodb

import (
	"gopkg.in/mgo.v2"
)

type ChallengeService struct {
	DB *DB
	Collection string
	database *mgo.Database
}

func (db *DB) NewChallengeService() *ChallengeService {
	svc := &ChallengeService{DB: db, Collection: "challenge", database: db.Session.DB(db.Settings.Database)}

	return svc
}
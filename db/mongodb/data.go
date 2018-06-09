package mongodb

import (
	"github.com/BeyondBankingDays/minions-api"
	"gopkg.in/mgo.v2"
	"log"
	"gopkg.in/mgo.v2/bson"
)

type DataService struct {
	DB         *DB
	Collection string
	database   *mgo.Database
}

func NewDataService(db *DB) *DataService {
	svc := &DataService{DB: db, Collection: "data", database: db.Session.DB(db.Settings.Database)}

	if err := svc.database.C(svc.Collection).EnsureIndexKey("time"); err != nil {
		log.Println("Could not create index", err)
	}
	if err := svc.database.C(svc.Collection).EnsureIndexKey("event"); err != nil {
		log.Println("Could not create index", err)
	}
	if err := svc.database.C(svc.Collection).EnsureIndexKey("type"); err != nil {
		log.Println("Could not create index", err)
	}
	if err := svc.database.C(svc.Collection).EnsureIndexKey("source.id"); err != nil {
		log.Println("Could not create index", err)
	}
	if err := svc.database.C(svc.Collection).EnsureIndexKey("source.type"); err != nil {
		log.Println("Could not create index", err)
	}
	return svc
}

func (s *DataService) CreateData(data *hackathon_api.Data) (*hackathon_api.Data, error) {
	data.Id = bson.NewObjectId()
	if err := s.database.C(s.Collection).Insert(data); err != nil {
		return nil, err
	}

	return data, nil
}

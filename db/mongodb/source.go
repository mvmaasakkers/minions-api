package mongodb

import (
	"github.com/BeyondBankingDays/minions-api"
	"gopkg.in/mgo.v2"
	"log"
	"gopkg.in/mgo.v2/bson"
	"errors"
)

type SourceService struct {
	DB *DB
	Collection string
	database *mgo.Database
}

func NewSourceService(db *DB) *SourceService {
	svc := &SourceService{DB: db, Collection: "source", database: db.Session.DB(db.Settings.Database)}

	if err := svc.database.C(svc.Collection).EnsureIndexKey("type"); err != nil {
		log.Println("Could not create index", err)
	}
	
	return svc
}

func (s *SourceService) CreateSource(source *hackathon_api.Source) (*hackathon_api.Source, error) {
	if source == nil {
		return nil, errors.New("cannot be nil")
	}
	source.Id = bson.NewObjectId()
	if err := s.database.C(s.Collection).Insert(source); err != nil {
		return nil, err
	}

	return source, nil
}

func (s *SourceService) EditSource(source *hackathon_api.Source) (*hackathon_api.Source, error) {
	if source == nil {
		return nil, errors.New("cannot be nil")
	}
	originalSource, err := s.GetSource(source.Id.Hex())
	if err != nil {
		return nil, err
	}

	if err := s.database.C(s.Collection).UpdateId(originalSource.Id, source); err != nil {
		return nil, err
	}

	return source, nil
}

func (s *SourceService) DeleteSource(id string) (error) {
	source, err := s.GetSource(id)
	if err != nil {
		return err
	}

	if err := s.database.C(s.Collection).RemoveId(source.Id); err != nil {
		return err
	}

	return nil
}

func (s *SourceService) GetSource(id string) (*hackathon_api.Source, error) {
	if !bson.IsObjectIdHex(id) {
		return nil, errors.New("invalid object id")
	}

	bsonId := bson.ObjectIdHex(id)
	source := &hackathon_api.Source{}
	if err := s.database.C(s.Collection).FindId(bsonId).One(source); err != nil {
		return nil, err
	}

	return source, nil
}

func (s *SourceService) ListSources() ([]*hackathon_api.Source, error) {
	sources := make([]*hackathon_api.Source, 0)
	if err := s.database.C(s.Collection).Find(bson.M{}).All(&sources); err != nil {
		return nil, err
	}

	return sources, nil
}

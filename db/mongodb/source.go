package mongodb

import (
	"github.com/jumba-nl/hackathon-api"
	"gopkg.in/mgo.v2"
	"log"
	"gopkg.in/mgo.v2/bson"
)

type SourceService struct {
	DB *DB
	Collection string
	database *mgo.Database
}

func (db *DB) NewSourceService() *SourceService {
	svc := &SourceService{DB: db, Collection: "source", database: db.Session.DB(db.Settings.Database)}

	if err := svc.database.C(svc.Collection).EnsureIndexKey("type"); err != nil {
		log.Println("Could not create index", err)
	}
	
	return svc
}

func (s *SourceService) CreateSource(source *hackathon_api.Source) (*hackathon_api.Source, error) {
	source.ID = bson.NewObjectId()
	if err := s.database.C(s.Collection).Insert(source); err != nil {
		return nil, err
	}

	return source, nil
}

func (s *SourceService) EditSource(source *hackathon_api.Source) (*hackathon_api.Source, error) {
	originalSource, err := s.GetSource(source.ID.Hex())
	if err != nil {
		return nil, err
	}

	if err := s.database.C(s.Collection).UpdateId(originalSource.ID, source); err != nil {
		return nil, err
	}

	return source, nil
}

func (s *SourceService) DeleteSource(id string) (error) {
	source, err := s.GetSource(id)
	if err != nil {
		return err
	}

	if err := s.database.C(s.Collection).RemoveId(source.ID); err != nil {
		return err
	}

	return nil
}

func (s *SourceService) GetSource(id string) (*hackathon_api.Source, error) {
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

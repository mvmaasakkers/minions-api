package mongodb

import (
	"github.com/BeyondBankingDays/minions-api"
	"gopkg.in/mgo.v2"
	"log"
	"gopkg.in/mgo.v2/bson"
	"time"
	"math/rand"
)

type TokenService struct {
	DB *DB
	Collection string
	database *mgo.Database
}

func (db *DB) NewTokenService() *TokenService {
	svc := &TokenService{DB: db, Collection: "token", database: db.Session.DB(db.Settings.Database)}

	if err := svc.database.C(svc.Collection).EnsureIndexKey("token"); err != nil {
		log.Println("Could not create index", err)
	}



	if err := svc.database.C(svc.Collection).EnsureIndexKey("expires_at"); err != nil {
		log.Println("Could not create index", err)
	}
	return svc
}

func (s *TokenService) CreateToken(token *hackathon_api.Token) (*hackathon_api.Token, error) {
	token.Id = bson.NewObjectId()
	token.CreatedAt = time.Now()
	token.ExpiresAt = time.Now().Add(time.Hour * 72)
	token.Token = generateToken()

	if err := s.database.C(s.Collection).Insert(token); err != nil {
		return nil, err
	}

	return token, nil
}

func (s *TokenService) DeleteToken(id string) (error) {
	token, err := s.GetToken(id)
	if err != nil {
		return err
	}

	if err := s.database.C(s.Collection).RemoveId(token.Id); err != nil {
		return err
	}

	return nil
}

func (s *TokenService) GetToken(id string) (*hackathon_api.Token, error) {
	bsonId := bson.ObjectIdHex(id)
	token := &hackathon_api.Token{}
	if err := s.database.C(s.Collection).FindId(bsonId).One(token); err != nil {
		return nil, err
	}

	return token, nil
}


func (s *TokenService) GetTokenByToken(tokenVal string) (*hackathon_api.Token, error) {
	token := &hackathon_api.Token{}
	if err := s.database.C(s.Collection).Find(bson.M{"token": tokenVal}).One(token); err != nil {
		return nil, err
	}

	return token, nil
}

func (s *TokenService) ListTokens() ([]*hackathon_api.Token, error) {
	tokens := make([]*hackathon_api.Token, 0)
	if err := s.database.C(s.Collection).Find(bson.M{}).All(&tokens); err != nil {
		return nil, err
	}

	return tokens, nil
}

// Generates a random string
func generateToken() string {
	length := 64
	rand.Seed(time.Now().UTC().UnixNano())

	var char string = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"

	buf := make([]byte, length)
	for i := 0; i < length; i++ {
		buf[i] = char[rand.Intn(len(char)-1)]
	}
	return string(buf)
}
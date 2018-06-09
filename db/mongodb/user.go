package mongodb

import (
	"github.com/BeyondBankingDays/minions-api"
	"gopkg.in/mgo.v2"
	"log"
	"gopkg.in/mgo.v2/bson"
	"math/rand"
	"time"
)

type UserService struct {
	DB *DB
	Collection string
	database *mgo.Database
}

func NewUserService(db *DB) *UserService {
	svc := &UserService{DB: db, Collection: "user", database: db.Session.DB(db.Settings.Database)}

	if err := svc.database.C(svc.Collection).EnsureIndexKey("username"); err != nil {
		log.Println("Could not create index", err)
	}

	return svc
}

func (s *UserService) CreateUser(user *hackathon_api.User) (*hackathon_api.User, error) {
	user.Id = bson.NewObjectId()
	if err := s.database.C(s.Collection).Insert(user); err != nil {
		return nil, err
	}

	return user, nil
}

func (s *UserService) EditUser(user *hackathon_api.User) (*hackathon_api.User, error) {
	originalUser, err := s.GetUser(user.Id.Hex())
	if err != nil {
		return nil, err
	}

	if err := s.database.C(s.Collection).UpdateId(originalUser.Id, user); err != nil {
		return nil, err
	}

	return user, nil
}

func (s *UserService) DeleteUser(id string) (error) {
	user, err := s.GetUser(id)
	if err != nil {
		return err
	}

	if err := s.database.C(s.Collection).RemoveId(user.Id); err != nil {
		return err
	}

	return nil
}

func (s *UserService) GetUser(id string) (*hackathon_api.User, error) {
	bsonId := bson.ObjectIdHex(id)
	user := &hackathon_api.User{}
	if err := s.database.C(s.Collection).FindId(bsonId).One(user); err != nil {
		return nil, err
	}

	if user.Score.Current == 0 {
		user.Score.Current = randomInt(1000, 99999)
		if err := s.database.C(s.Collection).UpdateId(user.Id, user); err != nil {
			return nil, err
		}
	}

	return user, nil
}


func randomInt(min int, max int) int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	p := r.Perm(max - min + 1)
	return p[min]
}

func (s *UserService) GetUserByUsername(username string) (*hackathon_api.User, error) {
	user := &hackathon_api.User{}
	if err := s.database.C(s.Collection).Find(bson.M{"username": username}).One(user); err != nil {
		return nil, err
	}

	return user, nil
}

func (s *UserService) ListUsers() ([]*hackathon_api.User, error) {
	users := make([]*hackathon_api.User, 0)
	if err := s.database.C(s.Collection).Find(bson.M{}).All(&users); err != nil {
		return nil, err
	}

	return users, nil
}

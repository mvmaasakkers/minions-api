package mongodb

import (
	"github.com/BeyondBankingDays/minions-api"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
	"errors"
)

type BankAccountService struct {
	DB *DB
	Collection string
	database *mgo.Database
}

func (db *DB) NewBankAccountService() *BankAccountService {
	svc := &BankAccountService{DB: db, Collection: "bank_account", database: db.Session.DB(db.Settings.Database)}

	if err := svc.database.C(svc.Collection).EnsureIndexKey("user_id"); err != nil {
		log.Println("Could not create index", err)
	}

	return svc
}

func (s *BankAccountService) CreateBankAccount(bankAccount *hackathon_api.BankAccount) (*hackathon_api.BankAccount, error) {
	if bankAccount == nil {
		return nil, errors.New("cannot be nil")
	}
	if _, err := s.database.C(s.Collection).UpsertId(bankAccount.Id, bankAccount); err != nil {
		return nil, err
	}

	return bankAccount, nil
}

func (s *BankAccountService) DeleteBankAccount(userId, id string) (error) {
	bankAccount, err := s.GetBankAccount(userId, id)
	if err != nil {
		return err
	}

	if err := s.database.C(s.Collection).RemoveId(bankAccount.Id); err != nil {
		return err
	}

	return nil
}

func (s *BankAccountService) GetBankAccount(userId, id string) (*hackathon_api.BankAccount, error) {
	//bsonId := bson.ObjectIdHex(id)
	bankAccount := &hackathon_api.BankAccount{}
	if err := s.database.C(s.Collection).Find(bson.M{"_id": id, "user_id": userId}).One(bankAccount); err != nil {
		return nil, err
	}

	return bankAccount, nil
}

func (s *BankAccountService) ListBankAccounts(userId string) ([]*hackathon_api.BankAccount, error) {
	bankAccounts := make([]*hackathon_api.BankAccount, 0)
	if err := s.database.C(s.Collection).Find(bson.M{"user_id": userId}).All(&bankAccounts); err != nil {
		return nil, err
	}

	return bankAccounts, nil
}

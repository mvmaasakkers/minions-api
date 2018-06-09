package mongodb

import (
	"github.com/BeyondBankingDays/minions-api"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
)

type BankTransactionService struct {
	DB *DB
	Collection string
	database *mgo.Database
}

func (db *DB) NewBankTransactionService() *BankTransactionService {
	svc := &BankTransactionService{DB: db, Collection: "bank_transaction", database: db.Session.DB(db.Settings.Database)}

	if err := svc.database.C(svc.Collection).EnsureIndexKey("user_id"); err != nil {
		log.Println("Could not create index", err)
	}

	if err := svc.database.C(svc.Collection).EnsureIndexKey("this_account.id"); err != nil {
		log.Println("Could not create index", err)
	}

	return svc
}

func (s *BankTransactionService) CreateBankTransaction(bankTransaction *hackathon_api.BankTransaction) (*hackathon_api.BankTransaction, error) {
	if _, err := s.database.C(s.Collection).UpsertId(bankTransaction.Id, bankTransaction); err != nil {
		return nil, err
	}

	return bankTransaction, nil
}

func (s *BankTransactionService) DeleteBankTransaction(userId, id string) (error) {
	bankTransaction, err := s.GetBankTransaction(userId, id)
	if err != nil {
		return err
	}

	if err := s.database.C(s.Collection).RemoveId(bankTransaction.Id); err != nil {
		return err
	}

	return nil
}

func (s *BankTransactionService) GetBankTransaction(userId, id string) (*hackathon_api.BankTransaction, error) {

	bankTransaction := &hackathon_api.BankTransaction{}
	if err := s.database.C(s.Collection).Find(bson.M{"user_id": userId, "_id": id}).One(bankTransaction); err != nil {
		return nil, err
	}

	return bankTransaction, nil
}

func (s *BankTransactionService) ListBankTransactions(userId string) ([]*hackathon_api.BankTransaction, error) {
	bankTransactions := make([]*hackathon_api.BankTransaction, 0)
	if err := s.database.C(s.Collection).Find(bson.M{"user_id": userId}).All(&bankTransactions); err != nil {
		return nil, err
	}

	return bankTransactions, nil
}


func (s *BankTransactionService) ListBankTransactionsByAccount(userId, accountId string) ([]*hackathon_api.BankTransaction, error) {
	bankTransactions := make([]*hackathon_api.BankTransaction, 0)
	if err := s.database.C(s.Collection).Find(bson.M{"user_id": userId, "this_account.id": accountId}).All(&bankTransactions); err != nil {
		return nil, err
	}

	return bankTransactions, nil
}

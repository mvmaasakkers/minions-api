package mongodb

import (
	"testing"
	"github.com/BeyondBankingDays/minions-api"
	"gopkg.in/mgo.v2/bson"
)

func TestBankAccountService_List(t *testing.T) {
	bankAccountService := db.NewBankAccountService()
	_, err := bankAccountService.ListBankAccounts("testuser")
	if err != nil {
		t.Errorf("Something went wrong executing List(): %s", err.Error())
		t.FailNow()
	}
}


func TestBankAccountService_CreateBankAccount(t *testing.T) {
	bankAccountService := db.NewBankAccountService()

	if _, err := bankAccountService.CreateBankAccount(nil); err == nil {
		t.Error("Exception was expected")
		t.FailNow()
	}

	bankAccount := &hackathon_api.BankAccount{}
	bankAccount.Label = "Test"
	bankAccount.UserID = "testuser"
	bankAccount.Id = bson.NewObjectId().Hex()

	if _, err := bankAccountService.CreateBankAccount(bankAccount); err != nil {
		t.Error("Exception was given and not expected", err.Error())
		t.FailNow()
	}

	if _, err := bankAccountService.GetBankAccount("testuser", bankAccount.Id); err != nil {
		t.Error("Exception was given and not expected", err.Error())
	}

}

func TestBankAccountService_DeleteBankAccount(t *testing.T) {
	bankAccountService := db.NewBankAccountService()
	err := bankAccountService.DeleteBankAccount("testuser", bankAccountFixtures["first"].Id)
	if err == nil {
		t.Error("Exception was expected")
		t.FailNow()
	}

	if _, err := bankAccountService.GetBankAccount("testuser", bankAccountFixtures["first"].Id); err == nil {
		t.Error("Exception was expected")
	}
}

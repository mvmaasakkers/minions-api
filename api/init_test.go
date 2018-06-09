package api

import (
	"io/ioutil"
	"os"
	"testing"

	"gopkg.in/mgo.v2/dbtest"
	"gopkg.in/mgo.v2"
	"log"
	"github.com/BeyondBankingDays/minions-api"
	"gopkg.in/mgo.v2/bson"
	"github.com/BeyondBankingDays/minions-api/db/mongodb"
)

var (
	// server runs a MongoDB server on a random port and storing data in temp dir.
	server dbtest.DBServer
	// testDb provides a disposable profile.DB implementation.
	session *mgo.Session

	db *mongodb.DB
)

// TestMain wraps all tests with the needed initialized mock DB and fixtures
//  - https://golang.org/pkg/testing/#hdr-Main
func TestMain(m *testing.M) {
	tempDir, err := ioutil.TempDir("", "testing")
	if err != nil {
		panic(err)
	}
	os.Exit(runTests(m, tempDir))
}

func runTests(m *testing.M, tempDir string) int {
	server.SetPath(tempDir)
	// Make sure we just "claim" a single session.
	session = server.Session()
	defer func() {
		// And close it when we're done, just before shutting the fixtures server down.
		session.Close()
		// So this can verify no sessions are leaked.
		server.Stop()
	}()

	db = &mongodb.DB{}
	db.Session = session
	db.Settings.Database = "api_test"

	mongodb.Conn = db
	installFixtures(db)

	return m.Run()
}

var sourceFixtures map[string]*hackathon_api.Source

var bankAccountFixtures = map[string]*hackathon_api.BankAccount{
	"first": {Id: "first", Label: "first", UserID: "tesuser"},
	"second": {Id: "second", Label: "second", UserID: "tesuser"},
}

var userId bson.ObjectId

func init() {
	userId = bson.NewObjectId()

	userFixtures = map[string]*hackathon_api.User{
		"first": {Id: userId, Username: "first"},
	}

	tokenFixtures = map[string]*hackathon_api.Token{
		"first": {Id: bson.NewObjectId(), UserId: userId.Hex(), Token: "first"},
	}

	sourceFixtures = map[string]*hackathon_api.Source{
		"first": {Id: bson.NewObjectId(), Name: "first"},
		"second": {Id: bson.NewObjectId(), Name: "second"},
	}
}


var userFixtures map[string]*hackathon_api.User

var tokenFixtures map[string]*hackathon_api.Token

func installFixtures(db *mongodb.DB) {
	sourceService := mongodb.NewSourceService(db)
	for _, item := range sourceFixtures {
		if _, err := sourceService.CreateSource(item); err != nil {
			log.Println(err)
		}
	}
	bankAccountService := mongodb.NewBankAccountService(db)
	for _, item := range bankAccountFixtures {
		if err := bankAccountService.DB.Session.DB(db.Settings.Database).C("bank_account").Insert(item); err != nil {
			log.Println(err)
		}
	}

	userService := mongodb.NewUserService(db)
	for _, item := range userFixtures {
		if _, err := userService.CreateUser(item); err != nil {
			log.Println(err)
		}
	}

	tokenService := mongodb.NewTokenService(db)
	for _, item := range tokenFixtures {
		if _, err := tokenService.CreateToken(item); err != nil {
			log.Println(err)
		}
	}
}
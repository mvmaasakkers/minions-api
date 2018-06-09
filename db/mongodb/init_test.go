package mongodb

import (
	"io/ioutil"
	"os"
	"testing"

	"gopkg.in/mgo.v2/dbtest"
	"gopkg.in/mgo.v2"
	"log"
	"github.com/BeyondBankingDays/minions-api"
	"gopkg.in/mgo.v2/bson"
)

var (
	// server runs a MongoDB server on a random port and storing data in temp dir.
	server dbtest.DBServer
	// testDb provides a disposable profile.DB implementation.
	session *mgo.Session

	db *DB
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
		// And close it when we're done, just before shutting the test server down.
		session.Close()
		// So this can verify no sessions are leaked.
		server.Stop()
	}()

	db = &DB{}
	db.Session = session
	db.Settings.Database = "api_test"

	installFixtures(db)

	return m.Run()
}

var sourceFixtures = map[string]*hackathon_api.Source{
	"first": {Id: bson.NewObjectId(), Name: "first"},
	"second": {Id: bson.NewObjectId(), Name: "second"},
}

var bankAccountFixtures = map[string]*hackathon_api.BankAccount{
	"first": {Id: "first", Label: "first", UserID: "tesuser"},
	"second": {Id: "second", Label: "second", UserID: "tesuser"},
}

func installFixtures(db *DB) {
	sourceService := NewSourceService(db)
	for _, item := range sourceFixtures {
		if _, err := sourceService.CreateSource(item); err != nil {
			log.Println(err)
		}
	}
	bankAccountService := NewBankAccountService(db)
	for _, item := range bankAccountFixtures {
		if err := bankAccountService.database.C("bank_account").Insert(item); err != nil {
			log.Println(err)
		}
	}

}
package mongodb

import (
	"testing"
	"github.com/BeyondBankingDays/minions-api"
)

func TestSourceService_List(t *testing.T) {
	sourceService := db.NewSourceService()
	sources, err := sourceService.ListSources()
	if err != nil {
		t.Errorf("Something went wrong executing List(): %s", err.Error())
		t.FailNow()
	}

	if len(sources) == 0 {
		t.Errorf("Expected more than 0 sources")
	}
}

func TestSourceService_Get(t *testing.T) {
	sourceService := db.NewSourceService()
	source, err := sourceService.GetSource(sourceFixtures["first"].Id.Hex())
	if err != nil {
		t.Errorf("Something went wrong executing List(): %s", err.Error())
		t.FailNow()
	}

	if source.Name != sourceFixtures["first"].Name {
		t.Errorf("Expected %s, got %s", sourceFixtures["first"].Name, source.Name)
	}
}

func TestSourceService_EditSource(t *testing.T) {
	sourceService := db.NewSourceService()
	_, err := sourceService.EditSource(nil)
	if err == nil {
		t.Error("Exception was expected")
		t.FailNow()
	}

	source, _ := sourceService.GetSource(sourceFixtures["first"].Id.Hex())
	source.Name = "Other name"
	if _, err := sourceService.EditSource(source); err != nil {
		t.Errorf("Something went wrong executing Edit(): %s", err.Error())
		t.FailNow()
	}

	source, _ = sourceService.GetSource(sourceFixtures["first"].Id.Hex())
	if source.Name == sourceFixtures["first"].Name {
		t.Error("Expected names to be different")
	}

}

func TestSourceService_CreateSource(t *testing.T) {
	sourceService := db.NewSourceService()

	if _, err := sourceService.CreateSource(nil); err == nil {
		t.Error("Exception was expected")
		t.FailNow()
	}

	source := &hackathon_api.Source{}
	source.Name = "Test"

	if _, err := sourceService.CreateSource(source); err != nil {
		t.Error("Exception was given and not expected", err.Error())
		t.FailNow()
	}

	if _, err := sourceService.GetSource(source.Id.Hex()); err != nil {
		t.Error("Exception was given and not expected", err.Error())
	}

}

func TestSourceService_DeleteSource(t *testing.T) {
	sourceService := db.NewSourceService()
	err := sourceService.DeleteSource("")
	if err == nil {
		t.Error("Exception was expected")
		t.FailNow()
	}

	if err := sourceService.DeleteSource(sourceFixtures["first"].Id.Hex()); err != nil {
		t.Error("Exception was given", err.Error())
		t.Fail()
	}

	if _, err := sourceService.GetSource(sourceFixtures["first"].Id.Hex()); err == nil {
		t.Error("Exception was expected")
	}
}

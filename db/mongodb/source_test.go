package mongodb

import "testing"

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

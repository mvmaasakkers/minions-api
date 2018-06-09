package api

import (
	"testing"
	"net/http"
	"net/http/httptest"
)

func TestMeta_ChallengeListHandler(t *testing.T) {
	meta := Meta{DB: *db}

	req, err := http.NewRequest("GET", "/v1/challenges", nil)
	if err != nil {
		t.Fatal(err)
	}

	ContextSet(req, "user", userFixtures["first"])
	rr := httptest.NewRecorder()

	meta.ChallengeListHandler(rr, req)
	status := rr.Code;
	if status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
}


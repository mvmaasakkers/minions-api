package api

import (
	"testing"
	"net/http"
	"net/http/httptest"
	"github.com/BeyondBankingDays/minions-api"
)

func TestMeta_GetUserHandler(t *testing.T) {
	var tests = []struct {
		User       *hackathon_api.User
		StatusCode int
	}{
		{
			User:       userFixtures["first"],
			StatusCode: http.StatusOK,
		},
		{
			User:       userFixtures["second"],
			StatusCode: http.StatusOK,
		},
		{
			User:       userFixtures["third"],
			StatusCode: http.StatusOK,
		},
		{
			User:       userFixtures["fourth"],
			StatusCode: http.StatusOK,
		},
		{
			User:       userFixtures["first"],
			StatusCode: http.StatusOK,
		},
		{
			User:       userFixtures["second"],
			StatusCode: http.StatusOK,
		},
		{
			User:       userFixtures["third"],
			StatusCode: http.StatusOK,
		},
		{
			User:       userFixtures["fourth"],
			StatusCode: http.StatusOK,
		},
		{
			User:       userFixtures["second"],
			StatusCode: http.StatusOK,
		},
		{
			User:       userFixtures["third"],
			StatusCode: http.StatusOK,
		},
		{
			User:       userFixtures["fourth"],
			StatusCode: http.StatusOK,
		},
		{
			User:       userFixtures["second"],
			StatusCode: http.StatusOK,
		},
		{
			User:       userFixtures["third"],
			StatusCode: http.StatusOK,
		},
		{
			User:       userFixtures["fourth"],
			StatusCode: http.StatusOK,
		},
		{
			User:       userFixtures["second"],
			StatusCode: http.StatusOK,
		},
		{
			User:       userFixtures["third"],
			StatusCode: http.StatusOK,
		},
		{
			User:       userFixtures["fourth"],
			StatusCode: http.StatusOK,
		},
		{
			User:       userFixtures["first"],
			StatusCode: http.StatusOK,
		},
		{
			User:       userFixtures["second"],
			StatusCode: http.StatusOK,
		},
		{
			User:       userFixtures["third"],
			StatusCode: http.StatusOK,
		},
		{
			User:       userFixtures["fourth"],
			StatusCode: http.StatusOK,
		},
	}
	meta := Meta{DB: *db}

	for _, test := range tests {
		req, err := http.NewRequest("GET", "/v1/user", nil)
		if err != nil {
			t.Fatal(err)
		}

		ContextSet(req, "user", test.User)
		rr := httptest.NewRecorder()

		meta.GetUserHandler(rr, req)
		status := rr.Code
		if test.StatusCode != status {
			t.Errorf("handler returned wrong status code: got %v want %v", status, test.StatusCode)
		}
	}

}

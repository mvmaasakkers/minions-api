package api

import (
	"testing"
	"net/http"
	"net/http/httptest"
)

func TestMeta_BankAccountListHandler(t *testing.T) {
	meta := Meta{DB: *db}

	req, err := http.NewRequest("GET", "/v1/bank/accounts", nil)
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


func TestMeta_BankTransactionsListHandler(t *testing.T) {
	meta := Meta{DB: *db}

	req, err := http.NewRequest("GET", "/v1/bank/transactions", nil)
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



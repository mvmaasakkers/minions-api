package api

import (
	"net/http"
	"github.com/BeyondBankingDays/minions-api/ext/bb"
	"github.com/BeyondBankingDays/minions-api"
	"log"
	"github.com/gorilla/mux"
	"github.com/BeyondBankingDays/minions-api/db/mongodb"
)

func (h *Meta) BankAccountListHandler(w http.ResponseWriter, r *http.Request) {
	user := getContextUser(r)
	bankAccountService := mongodb.NewBankAccountService(&h.DB)
	bankAccounts, err := bankAccountService.ListBankAccounts(user.Id.Hex())
	if err != nil {
		JsonResponse(w, r, http.StatusInternalServerError, NewApiError(err.Error()))
		return
	}

	JsonResponse(w, r, http.StatusOK, bankAccounts)
}


func (h *Meta) BankAccountGetHandler(w http.ResponseWriter, r *http.Request) {
	user := getContextUser(r)
	bankAccountService := mongodb.NewBankAccountService(&h.DB)
	vars := mux.Vars(r)
	if _, ok  := vars["id"]; !ok {
		JsonResponse(w, r, http.StatusBadRequest, NewApiError("no id given"))
		return
	}

	bankAccount, err := bankAccountService.GetBankAccount(user.Id.Hex(), vars["id"])
	if err != nil {
		JsonResponse(w, r, http.StatusInternalServerError, NewApiError(err.Error()))
		return
	}

	JsonResponse(w, r, http.StatusOK, bankAccount)
}

func (h *Meta) BankTransactionsListHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	if _, ok  := vars["id"]; !ok {
		JsonResponse(w, r, http.StatusBadRequest, NewApiError("no id given"))
		return
	}

	user := getContextUser(r)
	bankAccountService := mongodb.NewBankAccountService(&h.DB)
	bankAccount, err := bankAccountService.GetBankAccount(user.Id.Hex(), vars["id"])
	if err != nil {
		JsonResponse(w, r, http.StatusInternalServerError, NewApiError(err.Error()))
		return
	}

	bankTransactionsService := mongodb.NewBankTransactionService(&h.DB)
	transactions, err := bankTransactionsService.ListBankTransactionsByAccount(user.Id.Hex(), bankAccount.Id)
	if err != nil {
		JsonResponse(w, r, http.StatusInternalServerError, NewApiError(err.Error()))
		return
	}

	JsonResponse(w, r, http.StatusOK, transactions)
}

func (h *Meta) BankSyncData(w http.ResponseWriter, r *http.Request) {
	user := getContextUser(r)
	go h.bankSyncData(user)

	JsonResponse(w, r, http.StatusOK, Received{true, "started syncing bank data for user in background"})
}

func (h *Meta) bankSyncData(user *hackathon_api.User)  {
	for _, bankUser := range user.BankUsers {
		tokenResponse, err := bb.Login(bankUser)
		if err != nil {
			log.Println("Sync failed because there was a problem with login", err.Error())
			return
		}

		conn := bb.Conn{Token: tokenResponse.Token}
		accounts, err := conn.GetAccounts()
		if err != nil {
			log.Println("Sync failed because there was a problem with getting bankaccounts", err.Error())
			return
		}

		bankAccountService := mongodb.NewBankAccountService(&h.DB)
		bankTransactionService := mongodb.NewBankTransactionService(&h.DB)
		for _, account := range accounts.BankAccounts {
			account.UserID = user.Id.Hex()


			transactions, err := conn.GetTransactions(account)
			if err != nil {
				log.Println("Sync failed because there was a problem with getting transactions", err.Error())
				return
			}

			i := 0

			for _, transaction := range transactions.BankTransactions {
				transaction.UserID = user.Id.Hex()
				if _, err := bankTransactionService.CreateBankTransaction(&transaction); err != nil {
					log.Println("Error upserting transaction", err.Error())
				}
				i++
			}

			account.Transactions = i

			if _, err := bankAccountService.CreateBankAccount(&account); err != nil {
				log.Println("Error upserting bankaccount", err.Error())
			}
		}
	}
}
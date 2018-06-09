package api

import (
	"github.com/BeyondBankingDays/minions-api"
	"fmt"
	"net/http"
	"encoding/json"
)

const ConsumerKey = "diwtn1rinc4c1jxvtff5xsg3gxb42hofyhmpbts5"
const BaseURL = "https://beyondbanking.openbankproject.com"

type bbConn struct {
	Token string
}

func bbLogin(bankUser hackathon_api.BankUser) (*TokenResponse, error) {

	authorizationHeader := fmt.Sprintf(
		`DirectLogin username="%s",password="%s",consumer_key="%s"`,
		bankUser.Username,
		bankUser.Password,
		ConsumerKey,
	)

	client := http.Client{}

	req, err := http.NewRequest("POST", BaseURL+"/my/logins/direct", nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", authorizationHeader)
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	token := &TokenResponse{}
	if err := json.NewDecoder(resp.Body).Decode(token); err != nil {
		return nil, err
	}

	return token, nil
}

type TokenResponse struct {
	Token string `json:"token"`
}

func (c *bbConn) GetAccounts() (*hackathon_api.BankAccounts, error) {
	authHeader := fmt.Sprintf(`DirectLogin token="%s"`, c.Token)

	accounts := &hackathon_api.BankAccounts{}

	client := http.Client{}
	req, err := http.NewRequest("GET", BaseURL+"/obp/v3.0.0/my/accounts", nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", authHeader)

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	if err := json.NewDecoder(resp.Body).Decode(accounts); err != nil {
		return nil, err
	}

	return accounts, nil
}

func (c *bbConn) GetTransactions(account hackathon_api.BankAccount) (*hackathon_api.BankTransactions, error){
	authHeader := fmt.Sprintf(`DirectLogin token="%s"`, c.Token)
	url := fmt.Sprintf("/obp/v3.0.0/my/banks/%s/accounts/%s/transactions", account.BankID, account.Id)

	transactions := &hackathon_api.BankTransactions{}

	client := http.Client{}
	req, err := http.NewRequest("GET", BaseURL+url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", authHeader)

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	if err := json.NewDecoder(resp.Body).Decode(transactions); err != nil {
		return nil, err
	}

	return transactions, nil
}

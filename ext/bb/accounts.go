package bb

import (
	"net/http"
	"fmt"
	"encoding/json"
)

type Account struct {
	Id     string `json:"id"`
	Label  string `json:"label"`
	BankID string `json:"bank_id"`
	AccountRouting struct {
		Scheme  string `json:"scheme"`
		Address string `json:"address"`
	} `json:"account_routing"`
}

type Accounts struct {
	Accounts []Account `json:"accounts"`
}

func (c *Conn) GetAccounts() ([]Account, error) {
	authHeader := fmt.Sprintf(`DirectLogin token="%s"`, c.Token)

	accounts := Accounts{}

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
	if err := json.NewDecoder(resp.Body).Decode(&accounts); err != nil {
		return nil, err
	}

	return accounts.Accounts, nil
}

package bb

import (
	"net/http"
	"fmt"
	"encoding/json"
	"github.com/BeyondBankingDays/minions-api"
)


func (c *Conn) GetAccounts() (*hackathon_api.BankAccounts, error) {
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

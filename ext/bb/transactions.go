package bb

import (
	"fmt"
	"net/http"
	"encoding/json"
	"github.com/BeyondBankingDays/minions-api"
)

func (c *Conn) GetTransactions(account hackathon_api.BankAccount) (*hackathon_api.BankTransactions, error){
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
